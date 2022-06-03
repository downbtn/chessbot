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
func (b *Board) Display(outputFile string, size int) error {
	fp, err := os.Open(path.Join(AssetDir, "board.png"))
	defer fp.Close()
	if err != nil {
		return err
	}

	var pieceFiles = [...]string{"", "king_white.png", "queen_white.png", "pawn_white.png", "bishop_white.png", "knight_white.png", "rook_white.png", "king_black.png", "queen_black.png", "pawn_black.png", "bishop_black.png", "knight_black.png", "rook_black.png"}
	var pieceImgs [12]image.Image
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

	board := image.NewNRGBA(boardImg.Bounds())
	draw.Draw(board, boardImg.Bounds(), boardImg, image.ZP, draw.Src)
	for i, row := range b {
		for j, square := range row {
			if square == 0 {
				continue
			}
			sqBound := image.Rect(32*j, 32*i, 32*j+32, 32*i+32)
			pcImg := pieceImgs[square]
			draw.Draw(board, sqBound, pcImg, image.ZP, draw.Src)
		}
	}

	ofp, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	png.Encode(ofp, board)
	return nil
}
