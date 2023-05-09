package filter

import (
	"github.com/disintegration/imaging"
)

type Filter interface {
	Process(srcPath, dstPath string) error
}

type Grayscale struct{}

func (f *Grayscale) Process(srcPath, dstPath string) error {
	src, err := imaging.Open(srcPath)
	if err != nil {
		return err
	}
	img := imaging.Grayscale(src)
	img = imaging.AdjustContrast(img, 20)
	img = imaging.Sharpen(img, 2)

	// partie du prof
	// --------------
	// dstFile, err := os.Create(dstPath)
	// if err != nil {
	// 	return err
	// }
	// defer dstFile.Close()

	/* opts := jpeg.Options{Quality: 90}
	return jpeg.Encode (dstFile, dst, &opts*/
	// --------------

	// ma partie
	err = imaging.Save(img, dstPath)
	if err != nil {
		return err
	}

	return nil
}

type Blur struct{}

func (b *Blur) Process(srcPath, dstPath string) error {
	src, err := imaging.Open(srcPath)
	if err != nil {
		return err
	}
	img := imaging.Blur(src, 10)
	img = imaging.AdjustContrast(img, 20)
	img = imaging.Sharpen(img, 2)

	// partie du prof
	// --------------
	// dstFile, err := os.Create(dstPath)
	// if err != nil {
	// 	return err
	// }
	// defer dstFile.Close()

	/* opts := jpeg.Options{Quality: 90}
	return jpeg.Encode (dstFile, dst, &opts*/
	// --------------

	// ma partie
	err = imaging.Save(img, dstPath)
	if err != nil {
		return err
	}

	return nil
}
