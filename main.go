package main

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"mime"
	"os"
	"path"
	"text/template"

	"github.com/Luzifer/rconfig"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	log "github.com/sirupsen/logrus"
)

var (
	cfg = struct {
		BaseURL        string `flag:"base-url" default:"" description:"URL to prepend before filename" validate:"nonzero"`
		BasePath       string `flag:"base-path" default:"file/{{ printf \"%.2s\" .Hash }}/{{.Hash}}" description:"Path to upload the file to"`
		Bucket         string `flag:"bucket" default:"" description:"S3 bucket to upload files to" validate:"nonzero"`
		VersionAndExit bool   `flag:"version" default:"false" description:"Prints current version and exits"`
	}{}

	version = "dev"
)

func init() {
	if err := rconfig.ParseAndValidate(&cfg); err != nil {
		log.Fatalf("Unable to parse commandline options: %s", err)
	}

	if cfg.VersionAndExit {
		fmt.Printf("share %s\n", version)
		os.Exit(0)
	}
}

func main() {
	if len(rconfig.Args()) == 1 {
		log.Fatalf("Usage: share <file to upload>")
	}

	inFile := rconfig.Args()[1]
	upFile := path.Join(cfg.BasePath, path.Base(inFile))
	mimeType := mime.TypeByExtension(path.Ext(inFile))
	if mimeType == "" {
		mimeType = "application/octet-stream"
	}

	inFileHandle, err := os.Open(inFile)
	if err != nil {
		log.WithError(err).Fatal("Unable to open source file")
	}

	fileHash, err := hashFile(inFile)
	if err != nil {
		log.WithError(err).Fatal("Unable to calculate file hash")
	}
	upFile, err = executeTemplate(upFile, map[string]interface{}{
		"Hash": fileHash,
	})
	if err != nil {
		log.WithError(err).Fatal("Unable to assemble target path")
	}

	log.Infof("Uploading %q to %q with type %q", inFile, upFile, mimeType)

	sess := session.Must(session.NewSession())
	svc := s3.New(sess)

	if _, err := svc.PutObject(&s3.PutObjectInput{
		Body:        inFileHandle,
		Bucket:      aws.String(cfg.Bucket),
		ContentType: aws.String(mimeType),
		Key:         aws.String(upFile),
	}); err != nil {
		log.WithError(err).Fatalf("Unable to put file into S3")
	}

	fmt.Printf("%s%s", cfg.BaseURL, upFile)
}

func hashFile(inFile string) (string, error) {
	data, err := ioutil.ReadFile(inFile)
	if err != nil {
		return "", err
	}
	sum1 := sha1.Sum(data)
	return fmt.Sprintf("%x", sum1), nil
}

func executeTemplate(tplStr string, vars map[string]interface{}) (string, error) {
	tpl, err := template.New("basepath").Parse(tplStr)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	err = tpl.Execute(buf, vars)
	return buf.String(), err
}