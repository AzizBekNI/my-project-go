package tar

import (
	"archive/tar"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var filePath = "./file_uchun/"
var files = []string{filePath + "example.txt", "./vazifa.go"}

func addTar(filename string, tw *tar.Writer) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("fileni ochishda xatolik yuz berdi ! %v", err)
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		return fmt.Errorf("Filene malumotlarini olishda muamoga yuz berdi %v", err)

	}
	hdr := &tar.Header{
		Typeflag:   0,
		Name:       filename,
		Linkname:   "",
		Size:       stat.Size(),
		Mode:       int64(stat.Mode()),
		Uid:        0,
		Gid:        0,
		Uname:      "",
		Gname:      "",
		ModTime:    time.Time{},
		AccessTime: time.Time{},
		ChangeTime: time.Time{},
		Devmajor:   0,
		Devminor:   0,
		Xattrs:     map[string]string{},
		PAXRecords: map[string]string{},
		Format:     0,
	}
	if err = tw.WriteHeader(hdr); err != nil {
		return fmt.Errorf("Writerda xato bor !! %v", err)
	}
	copyed, err := io.Copy(tw, file)
	if err != nil {
		return fmt.Errorf("Fileni copy qilayotganda muammo bor !!")
	}
	if copyed < stat.Size() {
		return fmt.Errorf("file to'liqkigicha targa o'tmagan !! %v", err)
	}
	return nil
}
func ArchiveTar(archiveFilename string) int {
	if len(archiveFilename) == 0 {
		return -1
	}
	flags := os.O_WRONLY | os.O_TRUNC | os.O_CREATE
	file, err := os.OpenFile(archiveFilename, flags, 0644)
	if err != nil {
		fmt.Println("File ochishda mauamma bor !", err)
		return -1
	}
	defer file.Close()
	tw := tar.NewWriter(file)
	defer tw.Close()
	for _, fileName := range files {
		if err := addTar(fileName, tw); err != nil {
			log.Fatal("Fayllar Tar filega qo'shilayotganda muammo bor !!")
			return -1
		}

	}
	return 1
}
func Get_Tar() {
	err := ArchiveTar("first_archive.tar")
	if err != 1 {
		fmt.Println("faile targa o'tishda xatolik")
	} else {
		fmt.Println("fayl targa muafiqiyatli o'tdi")
	}
}
