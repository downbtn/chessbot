package chess

import (
	"image"
	"image/draw"
	"image/png"
	"os"
	"path"
)

const AssetDir string = "/home/downbtn/proj/chessbot/assets/"

// Display renders an image of the current state of the board to a given output file.
func (b *Board) Display(outputFile string) error {
	fp, err := os.Open(path.Join(AssetDir, "board.png"))
	defer fp.Close()
	if err != nil {
		return err
	}

	var pieceFiles = [...]string{"", "king_w.png", "queen_w.png", "pawn_w.png", "bishop_w.png", "knight_w.png", "rook_w.png", "king_b.png", "queen_b.png", "pawn_b.png", "bishop_b.png", "knight_b.png", "rook_b.png"}
	var pieceImgs [13]image.Image
	for i, p := range pieceFiles {
		if i == 0 {
			continue
		}

		pcf, err := os.Open(path.Join(AssetDir, p))
		defer pcf.Close()
		if err != nil {
			return err
		}

		pcImg, err := png.Decode(pcf)
		if err != nil {
			return err
		}
		pieceImgs[i] = pcImg
	}

	boardImg, err := png.Decode(fp)
	if err != nil {
		return err
	}

	// make a nrgba image then paste the board background on it
	board := image.NewNRGBA(boardImg.Bounds())
	draw.Draw(board, boardImg.Bounds(), boardImg, image.ZP, draw.Src)

	for i, row := range b {
		for j, square := range row {
			if square == 0 {
				continue
			}
			sqBound := image.Rect(34*j+2, 34*i+2, 34*j+34, 34*i+34)
			pcImg := pieceImgs[square]
			draw.Draw(board, sqBound, pcImg, image.ZP, draw.Over)
		}
	}

	ofp, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer ofp.Close()
	png.Encode(ofp, board)
	return nil
}
