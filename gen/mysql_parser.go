// Code generated from /Users/zhangzhiqiang/go/src/github.com/Aiyane/golinq/sql-parser/MySqlParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // MySqlParser

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa


var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 86, 622, 
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13, 
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9, 
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23, 
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4, 
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34, 
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9, 
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 3, 2, 3, 2, 3, 2, 3, 2, 5, 
	2, 89, 10, 2, 3, 3, 3, 3, 5, 3, 93, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 
	7, 4, 100, 10, 4, 12, 4, 14, 4, 103, 11, 4, 3, 5, 3, 5, 5, 5, 107, 10, 
	5, 3, 6, 3, 6, 7, 6, 111, 10, 6, 12, 6, 14, 6, 114, 11, 6, 3, 6, 3, 6, 
	3, 6, 7, 6, 119, 10, 6, 12, 6, 14, 6, 122, 11, 6, 3, 6, 3, 6, 5, 6, 126, 
	10, 6, 3, 7, 3, 7, 3, 7, 7, 7, 131, 10, 7, 12, 7, 14, 7, 134, 11, 7, 3, 
	8, 3, 8, 5, 8, 138, 10, 8, 3, 8, 5, 8, 141, 10, 8, 3, 8, 3, 8, 3, 8, 3, 
	8, 3, 8, 5, 8, 148, 10, 8, 3, 8, 5, 8, 151, 10, 8, 3, 8, 3, 8, 3, 8, 3, 
	8, 3, 8, 3, 8, 5, 8, 159, 10, 8, 3, 9, 5, 9, 162, 10, 9, 3, 9, 3, 9, 3, 
	9, 3, 9, 5, 9, 168, 10, 9, 3, 9, 3, 9, 5, 9, 172, 10, 9, 3, 9, 3, 9, 3, 
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 183, 10, 9, 5, 9, 185, 10, 
	9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 195, 
	10, 10, 3, 11, 3, 11, 5, 11, 199, 10, 11, 3, 11, 3, 11, 5, 11, 203, 10, 
	11, 3, 11, 5, 11, 206, 10, 11, 3, 11, 5, 11, 209, 10, 11, 3, 12, 3, 12, 
	5, 12, 213, 10, 12, 3, 12, 3, 12, 7, 12, 217, 10, 12, 12, 12, 14, 12, 220, 
	11, 12, 3, 13, 3, 13, 5, 13, 224, 10, 13, 3, 13, 5, 13, 227, 10, 13, 3, 
	13, 3, 13, 5, 13, 231, 10, 13, 3, 13, 5, 13, 234, 10, 13, 3, 13, 3, 13, 
	5, 13, 238, 10, 13, 3, 13, 5, 13, 241, 10, 13, 5, 13, 243, 10, 13, 3, 14, 
	3, 14, 3, 14, 3, 14, 5, 14, 249, 10, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 
	14, 7, 14, 256, 10, 14, 12, 14, 14, 14, 259, 11, 14, 3, 14, 3, 14, 5, 14, 
	263, 10, 14, 5, 14, 265, 10, 14, 3, 15, 3, 15, 5, 15, 269, 10, 15, 3, 16, 
	3, 16, 3, 16, 5, 16, 274, 10, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 280, 
	10, 16, 3, 17, 3, 17, 7, 17, 284, 10, 17, 12, 17, 14, 17, 287, 11, 17, 
	3, 18, 3, 18, 3, 19, 5, 19, 292, 10, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 
	20, 3, 20, 3, 20, 3, 20, 3, 20, 5, 20, 303, 10, 20, 3, 20, 5, 20, 306, 
	10, 20, 3, 21, 3, 21, 3, 21, 7, 21, 311, 10, 21, 12, 21, 14, 21, 314, 11, 
	21, 3, 22, 3, 22, 3, 22, 7, 22, 319, 10, 22, 12, 22, 14, 22, 322, 11, 22, 
	3, 23, 3, 23, 3, 23, 3, 23, 5, 23, 328, 10, 23, 3, 23, 5, 23, 331, 10, 
	23, 3, 24, 3, 24, 3, 24, 6, 24, 336, 10, 24, 13, 24, 14, 24, 337, 3, 24, 
	3, 24, 5, 24, 342, 10, 24, 3, 24, 3, 24, 3, 24, 3, 24, 6, 24, 348, 10, 
	24, 13, 24, 14, 24, 349, 3, 24, 3, 24, 5, 24, 354, 10, 24, 3, 24, 3, 24, 
	5, 24, 358, 10, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 
	26, 7, 26, 368, 10, 26, 12, 26, 14, 26, 371, 11, 26, 3, 27, 3, 27, 3, 27, 
	3, 27, 3, 27, 3, 27, 5, 27, 379, 10, 27, 3, 28, 3, 28, 3, 28, 3, 28, 5, 
	28, 385, 10, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 5, 29, 393, 
	10, 29, 3, 29, 3, 29, 3, 29, 5, 29, 398, 10, 29, 3, 29, 3, 29, 3, 29, 3, 
	29, 7, 29, 404, 10, 29, 12, 29, 14, 29, 407, 11, 29, 3, 30, 3, 30, 3, 30, 
	3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 5, 30, 418, 10, 30, 3, 30, 3, 
	30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 5, 30, 427, 10, 30, 3, 30, 3, 30, 
	3, 30, 3, 30, 5, 30, 433, 10, 30, 3, 30, 3, 30, 3, 30, 3, 30, 5, 30, 439, 
	10, 30, 3, 30, 3, 30, 3, 30, 3, 30, 5, 30, 445, 10, 30, 3, 30, 3, 30, 3, 
	30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 7, 30, 
	459, 10, 30, 12, 30, 14, 30, 462, 11, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 
	31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 7, 31, 475, 10, 31, 12, 31, 
	14, 31, 478, 11, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 
	31, 3, 31, 3, 31, 3, 31, 5, 31, 491, 10, 31, 3, 31, 3, 31, 3, 31, 3, 31, 
	3, 31, 3, 31, 3, 31, 7, 31, 500, 10, 31, 12, 31, 14, 31, 503, 11, 31, 3, 
	32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 
	3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 5, 33, 521, 10, 33, 3, 34, 3, 34, 3, 
	35, 3, 35, 3, 36, 3, 36, 5, 36, 529, 10, 36, 3, 36, 5, 36, 532, 10, 36, 
	3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 5, 36, 539, 10, 36, 3, 36, 3, 36, 3, 
	37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 7, 37, 
	553, 10, 37, 12, 37, 14, 37, 556, 11, 37, 5, 37, 558, 10, 37, 3, 38, 3, 
	38, 3, 38, 3, 38, 3, 39, 3, 39, 3, 39, 7, 39, 567, 10, 39, 12, 39, 14, 
	39, 570, 11, 39, 3, 40, 3, 40, 5, 40, 574, 10, 40, 3, 41, 3, 41, 5, 41, 
	578, 10, 41, 3, 41, 3, 41, 3, 41, 3, 41, 5, 41, 584, 10, 41, 3, 41, 5, 
	41, 587, 10, 41, 3, 41, 5, 41, 590, 10, 41, 3, 42, 3, 42, 5, 42, 594, 10, 
	42, 3, 42, 3, 42, 5, 42, 598, 10, 42, 3, 42, 5, 42, 601, 10, 42, 3, 42, 
	3, 42, 3, 42, 3, 42, 7, 42, 607, 10, 42, 12, 42, 14, 42, 610, 11, 42, 3, 
	42, 3, 42, 5, 42, 614, 10, 42, 3, 42, 5, 42, 617, 10, 42, 3, 42, 5, 42, 
	620, 10, 42, 3, 42, 2, 5, 56, 58, 60, 43, 2, 4, 6, 8, 10, 12, 14, 16, 18, 
	20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 
	56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 2, 16, 4, 2, 10, 
	10, 15, 15, 4, 2, 14, 14, 24, 24, 4, 2, 28, 28, 38, 38, 4, 2, 19, 19, 42, 
	42, 4, 2, 32, 32, 83, 83, 4, 2, 31, 31, 71, 71, 5, 2, 19, 19, 42, 42, 59, 
	59, 4, 2, 37, 37, 39, 39, 5, 2, 7, 7, 55, 55, 58, 58, 4, 2, 60, 62, 66, 
	67, 6, 2, 31, 31, 63, 63, 65, 65, 71, 72, 4, 2, 8, 8, 34, 34, 3, 2, 63, 
	65, 4, 2, 45, 45, 52, 52, 2, 699, 2, 88, 3, 2, 2, 2, 4, 92, 3, 2, 2, 2, 
	6, 94, 3, 2, 2, 2, 8, 104, 3, 2, 2, 2, 10, 125, 3, 2, 2, 2, 12, 127, 3, 
	2, 2, 2, 14, 158, 3, 2, 2, 2, 16, 184, 3, 2, 2, 2, 18, 194, 3, 2, 2, 2, 
	20, 196, 3, 2, 2, 2, 22, 212, 3, 2, 2, 2, 24, 242, 3, 2, 2, 2, 26, 244, 
	3, 2, 2, 2, 28, 266, 3, 2, 2, 2, 30, 270, 3, 2, 2, 2, 32, 281, 3, 2, 2, 
	2, 34, 288, 3, 2, 2, 2, 36, 291, 3, 2, 2, 2, 38, 305, 3, 2, 2, 2, 40, 307, 
	3, 2, 2, 2, 42, 315, 3, 2, 2, 2, 44, 330, 3, 2, 2, 2, 46, 357, 3, 2, 2, 
	2, 48, 359, 3, 2, 2, 2, 50, 364, 3, 2, 2, 2, 52, 378, 3, 2, 2, 2, 54, 384, 
	3, 2, 2, 2, 56, 397, 3, 2, 2, 2, 58, 408, 3, 2, 2, 2, 60, 490, 3, 2, 2, 
	2, 62, 504, 3, 2, 2, 2, 64, 520, 3, 2, 2, 2, 66, 522, 3, 2, 2, 2, 68, 524, 
	3, 2, 2, 2, 70, 526, 3, 2, 2, 2, 72, 557, 3, 2, 2, 2, 74, 559, 3, 2, 2, 
	2, 76, 563, 3, 2, 2, 2, 78, 573, 3, 2, 2, 2, 80, 575, 3, 2, 2, 2, 82, 591, 
	3, 2, 2, 2, 84, 89, 5, 20, 11, 2, 85, 89, 5, 70, 36, 2, 86, 89, 5, 82, 
	42, 2, 87, 89, 5, 80, 41, 2, 88, 84, 3, 2, 2, 2, 88, 85, 3, 2, 2, 2, 88, 
	86, 3, 2, 2, 2, 88, 87, 3, 2, 2, 2, 89, 3, 3, 2, 2, 2, 90, 93, 5, 20, 11, 
	2, 91, 93, 5, 18, 10, 2, 92, 90, 3, 2, 2, 2, 92, 91, 3, 2, 2, 2, 93, 5, 
	3, 2, 2, 2, 94, 95, 7, 35, 2, 2, 95, 96, 7, 12, 2, 2, 96, 101, 5, 8, 5, 
	2, 97, 98, 7, 79, 2, 2, 98, 100, 5, 8, 5, 2, 99, 97, 3, 2, 2, 2, 100, 103, 
	3, 2, 2, 2, 101, 99, 3, 2, 2, 2, 101, 102, 3, 2, 2, 2, 102, 7, 3, 2, 2, 
	2, 103, 101, 3, 2, 2, 2, 104, 106, 5, 56, 29, 2, 105, 107, 9, 2, 2, 2, 
	106, 105, 3, 2, 2, 2, 106, 107, 3, 2, 2, 2, 107, 9, 3, 2, 2, 2, 108, 112, 
	5, 14, 8, 2, 109, 111, 5, 16, 9, 2, 110, 109, 3, 2, 2, 2, 111, 114, 3, 
	2, 2, 2, 112, 110, 3, 2, 2, 2, 112, 113, 3, 2, 2, 2, 113, 126, 3, 2, 2, 
	2, 114, 112, 3, 2, 2, 2, 115, 116, 7, 77, 2, 2, 116, 120, 5, 14, 8, 2, 
	117, 119, 5, 16, 9, 2, 118, 117, 3, 2, 2, 2, 119, 122, 3, 2, 2, 2, 120, 
	118, 3, 2, 2, 2, 120, 121, 3, 2, 2, 2, 121, 123, 3, 2, 2, 2, 122, 120, 
	3, 2, 2, 2, 123, 124, 7, 78, 2, 2, 124, 126, 3, 2, 2, 2, 125, 108, 3, 2, 
	2, 2, 125, 115, 3, 2, 2, 2, 126, 11, 3, 2, 2, 2, 127, 132, 5, 10, 6, 2, 
	128, 129, 7, 79, 2, 2, 129, 131, 5, 10, 6, 2, 130, 128, 3, 2, 2, 2, 131, 
	134, 3, 2, 2, 2, 132, 130, 3, 2, 2, 2, 132, 133, 3, 2, 2, 2, 133, 13, 3, 
	2, 2, 2, 134, 132, 3, 2, 2, 2, 135, 140, 5, 32, 17, 2, 136, 138, 7, 9, 
	2, 2, 137, 136, 3, 2, 2, 2, 137, 138, 3, 2, 2, 2, 138, 139, 3, 2, 2, 2, 
	139, 141, 7, 85, 2, 2, 140, 137, 3, 2, 2, 2, 140, 141, 3, 2, 2, 2, 141, 
	159, 3, 2, 2, 2, 142, 148, 5, 4, 3, 2, 143, 144, 7, 77, 2, 2, 144, 145, 
	5, 4, 3, 2, 145, 146, 7, 78, 2, 2, 146, 148, 3, 2, 2, 2, 147, 142, 3, 2, 
	2, 2, 147, 143, 3, 2, 2, 2, 148, 150, 3, 2, 2, 2, 149, 151, 7, 9, 2, 2, 
	150, 149, 3, 2, 2, 2, 150, 151, 3, 2, 2, 2, 151, 152, 3, 2, 2, 2, 152, 
	153, 7, 85, 2, 2, 153, 159, 3, 2, 2, 2, 154, 155, 7, 77, 2, 2, 155, 156, 
	5, 12, 7, 2, 156, 157, 7, 78, 2, 2, 157, 159, 3, 2, 2, 2, 158, 135, 3, 
	2, 2, 2, 158, 147, 3, 2, 2, 2, 158, 154, 3, 2, 2, 2, 159, 15, 3, 2, 2, 
	2, 160, 162, 9, 3, 2, 2, 161, 160, 3, 2, 2, 2, 161, 162, 3, 2, 2, 2, 162, 
	163, 3, 2, 2, 2, 163, 164, 7, 27, 2, 2, 164, 167, 5, 14, 8, 2, 165, 166, 
	7, 33, 2, 2, 166, 168, 5, 56, 29, 2, 167, 165, 3, 2, 2, 2, 167, 168, 3, 
	2, 2, 2, 168, 185, 3, 2, 2, 2, 169, 171, 9, 4, 2, 2, 170, 172, 7, 36, 2, 
	2, 171, 170, 3, 2, 2, 2, 171, 172, 3, 2, 2, 2, 172, 173, 3, 2, 2, 2, 173, 
	174, 7, 27, 2, 2, 174, 182, 5, 14, 8, 2, 175, 176, 7, 33, 2, 2, 176, 183, 
	5, 56, 29, 2, 177, 178, 7, 44, 2, 2, 178, 179, 7, 77, 2, 2, 179, 180, 5, 
	40, 21, 2, 180, 181, 7, 78, 2, 2, 181, 183, 3, 2, 2, 2, 182, 175, 3, 2, 
	2, 2, 182, 177, 3, 2, 2, 2, 183, 185, 3, 2, 2, 2, 184, 161, 3, 2, 2, 2, 
	184, 169, 3, 2, 2, 2, 185, 17, 3, 2, 2, 2, 186, 187, 7, 77, 2, 2, 187, 
	188, 5, 20, 11, 2, 188, 189, 7, 78, 2, 2, 189, 195, 3, 2, 2, 2, 190, 191, 
	7, 77, 2, 2, 191, 192, 5, 18, 10, 2, 192, 193, 7, 78, 2, 2, 193, 195, 3, 
	2, 2, 2, 194, 186, 3, 2, 2, 2, 194, 190, 3, 2, 2, 2, 195, 19, 3, 2, 2, 
	2, 196, 198, 7, 40, 2, 2, 197, 199, 7, 16, 2, 2, 198, 197, 3, 2, 2, 2, 
	198, 199, 3, 2, 2, 2, 199, 200, 3, 2, 2, 2, 200, 202, 5, 22, 12, 2, 201, 
	203, 5, 26, 14, 2, 202, 201, 3, 2, 2, 2, 202, 203, 3, 2, 2, 2, 203, 205, 
	3, 2, 2, 2, 204, 206, 5, 6, 4, 2, 205, 204, 3, 2, 2, 2, 205, 206, 3, 2, 
	2, 2, 206, 208, 3, 2, 2, 2, 207, 209, 5, 30, 16, 2, 208, 207, 3, 2, 2, 
	2, 208, 209, 3, 2, 2, 2, 209, 21, 3, 2, 2, 2, 210, 213, 7, 60, 2, 2, 211, 
	213, 5, 24, 13, 2, 212, 210, 3, 2, 2, 2, 212, 211, 3, 2, 2, 2, 213, 218, 
	3, 2, 2, 2, 214, 215, 7, 79, 2, 2, 215, 217, 5, 24, 13, 2, 216, 214, 3, 
	2, 2, 2, 217, 220, 3, 2, 2, 2, 218, 216, 3, 2, 2, 2, 218, 219, 3, 2, 2, 
	2, 219, 23, 3, 2, 2, 2, 220, 218, 3, 2, 2, 2, 221, 226, 5, 32, 17, 2, 222, 
	224, 7, 9, 2, 2, 223, 222, 3, 2, 2, 2, 223, 224, 3, 2, 2, 2, 224, 225, 
	3, 2, 2, 2, 225, 227, 7, 85, 2, 2, 226, 223, 3, 2, 2, 2, 226, 227, 3, 2, 
	2, 2, 227, 243, 3, 2, 2, 2, 228, 233, 5, 44, 23, 2, 229, 231, 7, 9, 2, 
	2, 230, 229, 3, 2, 2, 2, 230, 231, 3, 2, 2, 2, 231, 232, 3, 2, 2, 2, 232, 
	234, 7, 85, 2, 2, 233, 230, 3, 2, 2, 2, 233, 234, 3, 2, 2, 2, 234, 243, 
	3, 2, 2, 2, 235, 240, 5, 56, 29, 2, 236, 238, 7, 9, 2, 2, 237, 236, 3, 
	2, 2, 2, 237, 238, 3, 2, 2, 2, 238, 239, 3, 2, 2, 2, 239, 241, 7, 85, 2, 
	2, 240, 237, 3, 2, 2, 2, 240, 241, 3, 2, 2, 2, 241, 243, 3, 2, 2, 2, 242, 
	221, 3, 2, 2, 2, 242, 228, 3, 2, 2, 2, 242, 235, 3, 2, 2, 2, 243, 25, 3, 
	2, 2, 2, 244, 245, 7, 20, 2, 2, 245, 248, 5, 12, 7, 2, 246, 247, 7, 47, 
	2, 2, 247, 249, 5, 56, 29, 2, 248, 246, 3, 2, 2, 2, 248, 249, 3, 2, 2, 
	2, 249, 264, 3, 2, 2, 2, 250, 251, 7, 21, 2, 2, 251, 252, 7, 12, 2, 2, 
	252, 257, 5, 28, 15, 2, 253, 254, 7, 79, 2, 2, 254, 256, 5, 28, 15, 2, 
	255, 253, 3, 2, 2, 2, 256, 259, 3, 2, 2, 2, 257, 255, 3, 2, 2, 2, 257, 
	258, 3, 2, 2, 2, 258, 262, 3, 2, 2, 2, 259, 257, 3, 2, 2, 2, 260, 261, 
	7, 22, 2, 2, 261, 263, 5, 56, 29, 2, 262, 260, 3, 2, 2, 2, 262, 263, 3, 
	2, 2, 2, 263, 265, 3, 2, 2, 2, 264, 250, 3, 2, 2, 2, 264, 265, 3, 2, 2, 
	2, 265, 27, 3, 2, 2, 2, 266, 268, 5, 56, 29, 2, 267, 269, 9, 2, 2, 2, 268, 
	267, 3, 2, 2, 2, 268, 269, 3, 2, 2, 2, 269, 29, 3, 2, 2, 2, 270, 279, 7, 
	30, 2, 2, 271, 272, 7, 81, 2, 2, 272, 274, 7, 79, 2, 2, 273, 271, 3, 2, 
	2, 2, 273, 274, 3, 2, 2, 2, 274, 275, 3, 2, 2, 2, 275, 280, 7, 81, 2, 2, 
	276, 277, 7, 81, 2, 2, 277, 278, 7, 57, 2, 2, 278, 280, 7, 81, 2, 2, 279, 
	273, 3, 2, 2, 2, 279, 276, 3, 2, 2, 2, 280, 31, 3, 2, 2, 2, 281, 285, 7, 
	85, 2, 2, 282, 284, 7, 84, 2, 2, 283, 282, 3, 2, 2, 2, 284, 287, 3, 2, 
	2, 2, 285, 283, 3, 2, 2, 2, 285, 286, 3, 2, 2, 2, 286, 33, 3, 2, 2, 2, 
	287, 285, 3, 2, 2, 2, 288, 289, 9, 5, 2, 2, 289, 35, 3, 2, 2, 2, 290, 292, 
	7, 31, 2, 2, 291, 290, 3, 2, 2, 2, 291, 292, 3, 2, 2, 2, 292, 293, 3, 2, 
	2, 2, 293, 294, 9, 6, 2, 2, 294, 37, 3, 2, 2, 2, 295, 306, 7, 80, 2, 2, 
	296, 306, 7, 81, 2, 2, 297, 298, 7, 65, 2, 2, 298, 306, 7, 81, 2, 2, 299, 
	306, 5, 34, 18, 2, 300, 306, 7, 82, 2, 2, 301, 303, 7, 31, 2, 2, 302, 301, 
	3, 2, 2, 2, 302, 303, 3, 2, 2, 2, 303, 304, 3, 2, 2, 2, 304, 306, 9, 6, 
	2, 2, 305, 295, 3, 2, 2, 2, 305, 296, 3, 2, 2, 2, 305, 297, 3, 2, 2, 2, 
	305, 299, 3, 2, 2, 2, 305, 300, 3, 2, 2, 2, 305, 302, 3, 2, 2, 2, 306, 
	39, 3, 2, 2, 2, 307, 312, 7, 85, 2, 2, 308, 309, 7, 79, 2, 2, 309, 311, 
	7, 85, 2, 2, 310, 308, 3, 2, 2, 2, 311, 314, 3, 2, 2, 2, 312, 310, 3, 2, 
	2, 2, 312, 313, 3, 2, 2, 2, 313, 41, 3, 2, 2, 2, 314, 312, 3, 2, 2, 2, 
	315, 320, 5, 56, 29, 2, 316, 317, 7, 79, 2, 2, 317, 319, 5, 56, 29, 2, 
	318, 316, 3, 2, 2, 2, 319, 322, 3, 2, 2, 2, 320, 318, 3, 2, 2, 2, 320, 
	321, 3, 2, 2, 2, 321, 43, 3, 2, 2, 2, 322, 320, 3, 2, 2, 2, 323, 331, 5, 
	46, 24, 2, 324, 325, 7, 85, 2, 2, 325, 327, 7, 77, 2, 2, 326, 328, 5, 50, 
	26, 2, 327, 326, 3, 2, 2, 2, 327, 328, 3, 2, 2, 2, 328, 329, 3, 2, 2, 2, 
	329, 331, 7, 78, 2, 2, 330, 323, 3, 2, 2, 2, 330, 324, 3, 2, 2, 2, 331, 
	45, 3, 2, 2, 2, 332, 333, 7, 13, 2, 2, 333, 335, 5, 32, 17, 2, 334, 336, 
	5, 48, 25, 2, 335, 334, 3, 2, 2, 2, 336, 337, 3, 2, 2, 2, 337, 335, 3, 
	2, 2, 2, 337, 338, 3, 2, 2, 2, 338, 341, 3, 2, 2, 2, 339, 340, 7, 17, 2, 
	2, 340, 342, 5, 54, 28, 2, 341, 339, 3, 2, 2, 2, 341, 342, 3, 2, 2, 2, 
	342, 343, 3, 2, 2, 2, 343, 344, 7, 56, 2, 2, 344, 358, 3, 2, 2, 2, 345, 
	347, 7, 13, 2, 2, 346, 348, 5, 48, 25, 2, 347, 346, 3, 2, 2, 2, 348, 349, 
	3, 2, 2, 2, 349, 347, 3, 2, 2, 2, 349, 350, 3, 2, 2, 2, 350, 353, 3, 2, 
	2, 2, 351, 352, 7, 17, 2, 2, 352, 354, 5, 54, 28, 2, 353, 351, 3, 2, 2, 
	2, 353, 354, 3, 2, 2, 2, 354, 355, 3, 2, 2, 2, 355, 356, 7, 56, 2, 2, 356, 
	358, 3, 2, 2, 2, 357, 332, 3, 2, 2, 2, 357, 345, 3, 2, 2, 2, 358, 47, 3, 
	2, 2, 2, 359, 360, 7, 46, 2, 2, 360, 361, 5, 54, 28, 2, 361, 362, 7, 41, 
	2, 2, 362, 363, 5, 54, 28, 2, 363, 49, 3, 2, 2, 2, 364, 369, 5, 52, 27, 
	2, 365, 366, 7, 79, 2, 2, 366, 368, 5, 52, 27, 2, 367, 365, 3, 2, 2, 2, 
	368, 371, 3, 2, 2, 2, 369, 367, 3, 2, 2, 2, 369, 370, 3, 2, 2, 2, 370, 
	51, 3, 2, 2, 2, 371, 369, 3, 2, 2, 2, 372, 379, 7, 60, 2, 2, 373, 379, 
	7, 7, 2, 2, 374, 379, 5, 38, 20, 2, 375, 379, 5, 32, 17, 2, 376, 379, 5, 
	44, 23, 2, 377, 379, 5, 56, 29, 2, 378, 372, 3, 2, 2, 2, 378, 373, 3, 2, 
	2, 2, 378, 374, 3, 2, 2, 2, 378, 375, 3, 2, 2, 2, 378, 376, 3, 2, 2, 2, 
	378, 377, 3, 2, 2, 2, 379, 53, 3, 2, 2, 2, 380, 385, 5, 38, 20, 2, 381, 
	385, 5, 32, 17, 2, 382, 385, 5, 44, 23, 2, 383, 385, 5, 56, 29, 2, 384, 
	380, 3, 2, 2, 2, 384, 381, 3, 2, 2, 2, 384, 382, 3, 2, 2, 2, 384, 383, 
	3, 2, 2, 2, 385, 55, 3, 2, 2, 2, 386, 387, 8, 29, 1, 2, 387, 388, 9, 7, 
	2, 2, 388, 398, 5, 56, 29, 6, 389, 390, 5, 58, 30, 2, 390, 392, 7, 26, 
	2, 2, 391, 393, 7, 31, 2, 2, 392, 391, 3, 2, 2, 2, 392, 393, 3, 2, 2, 2, 
	393, 394, 3, 2, 2, 2, 394, 395, 9, 8, 2, 2, 395, 398, 3, 2, 2, 2, 396, 
	398, 5, 58, 30, 2, 397, 386, 3, 2, 2, 2, 397, 389, 3, 2, 2, 2, 397, 396, 
	3, 2, 2, 2, 398, 405, 3, 2, 2, 2, 399, 400, 12, 5, 2, 2, 400, 401, 5, 66, 
	34, 2, 401, 402, 5, 56, 29, 6, 402, 404, 3, 2, 2, 2, 403, 399, 3, 2, 2, 
	2, 404, 407, 3, 2, 2, 2, 405, 403, 3, 2, 2, 2, 405, 406, 3, 2, 2, 2, 406, 
	57, 3, 2, 2, 2, 407, 405, 3, 2, 2, 2, 408, 409, 8, 30, 1, 2, 409, 410, 
	5, 60, 31, 2, 410, 460, 3, 2, 2, 2, 411, 412, 12, 8, 2, 2, 412, 413, 5, 
	64, 33, 2, 413, 414, 5, 58, 30, 9, 414, 459, 3, 2, 2, 2, 415, 417, 12, 
	6, 2, 2, 416, 418, 7, 31, 2, 2, 417, 416, 3, 2, 2, 2, 417, 418, 3, 2, 2, 
	2, 418, 419, 3, 2, 2, 2, 419, 420, 7, 11, 2, 2, 420, 421, 5, 58, 30, 2, 
	421, 422, 7, 8, 2, 2, 422, 423, 5, 58, 30, 7, 423, 459, 3, 2, 2, 2, 424, 
	426, 12, 5, 2, 2, 425, 427, 7, 31, 2, 2, 426, 425, 3, 2, 2, 2, 426, 427, 
	3, 2, 2, 2, 427, 428, 3, 2, 2, 2, 428, 429, 7, 29, 2, 2, 429, 459, 5, 58, 
	30, 6, 430, 432, 12, 4, 2, 2, 431, 433, 7, 31, 2, 2, 432, 431, 3, 2, 2, 
	2, 432, 433, 3, 2, 2, 2, 433, 434, 3, 2, 2, 2, 434, 435, 9, 9, 2, 2, 435, 
	459, 5, 58, 30, 5, 436, 438, 12, 10, 2, 2, 437, 439, 7, 31, 2, 2, 438, 
	437, 3, 2, 2, 2, 438, 439, 3, 2, 2, 2, 439, 440, 3, 2, 2, 2, 440, 441, 
	7, 23, 2, 2, 441, 444, 7, 77, 2, 2, 442, 445, 5, 4, 3, 2, 443, 445, 5, 
	42, 22, 2, 444, 442, 3, 2, 2, 2, 444, 443, 3, 2, 2, 2, 445, 446, 3, 2, 
	2, 2, 446, 447, 7, 78, 2, 2, 447, 459, 3, 2, 2, 2, 448, 449, 12, 9, 2, 
	2, 449, 450, 7, 26, 2, 2, 450, 459, 5, 36, 19, 2, 451, 452, 12, 7, 2, 2, 
	452, 453, 5, 64, 33, 2, 453, 454, 9, 10, 2, 2, 454, 455, 7, 77, 2, 2, 455, 
	456, 5, 4, 3, 2, 456, 457, 7, 78, 2, 2, 457, 459, 3, 2, 2, 2, 458, 411, 
	3, 2, 2, 2, 458, 415, 3, 2, 2, 2, 458, 424, 3, 2, 2, 2, 458, 430, 3, 2, 
	2, 2, 458, 436, 3, 2, 2, 2, 458, 448, 3, 2, 2, 2, 458, 451, 3, 2, 2, 2, 
	459, 462, 3, 2, 2, 2, 460, 458, 3, 2, 2, 2, 460, 461, 3, 2, 2, 2, 461, 
	59, 3, 2, 2, 2, 462, 460, 3, 2, 2, 2, 463, 464, 8, 31, 1, 2, 464, 491, 
	5, 38, 20, 2, 465, 491, 5, 32, 17, 2, 466, 491, 5, 44, 23, 2, 467, 468, 
	5, 62, 32, 2, 468, 469, 5, 60, 31, 8, 469, 491, 3, 2, 2, 2, 470, 471, 7, 
	77, 2, 2, 471, 476, 5, 56, 29, 2, 472, 473, 7, 79, 2, 2, 473, 475, 5, 56, 
	29, 2, 474, 472, 3, 2, 2, 2, 475, 478, 3, 2, 2, 2, 476, 474, 3, 2, 2, 2, 
	476, 477, 3, 2, 2, 2, 477, 479, 3, 2, 2, 2, 478, 476, 3, 2, 2, 2, 479, 
	480, 7, 78, 2, 2, 480, 491, 3, 2, 2, 2, 481, 482, 7, 18, 2, 2, 482, 483, 
	7, 77, 2, 2, 483, 484, 5, 4, 3, 2, 484, 485, 7, 78, 2, 2, 485, 491, 3, 
	2, 2, 2, 486, 487, 7, 77, 2, 2, 487, 488, 5, 4, 3, 2, 488, 489, 7, 78, 
	2, 2, 489, 491, 3, 2, 2, 2, 490, 463, 3, 2, 2, 2, 490, 465, 3, 2, 2, 2, 
	490, 466, 3, 2, 2, 2, 490, 467, 3, 2, 2, 2, 490, 470, 3, 2, 2, 2, 490, 
	481, 3, 2, 2, 2, 490, 486, 3, 2, 2, 2, 491, 501, 3, 2, 2, 2, 492, 493, 
	12, 4, 2, 2, 493, 494, 9, 11, 2, 2, 494, 500, 5, 60, 31, 5, 495, 496, 12, 
	3, 2, 2, 496, 497, 5, 68, 35, 2, 497, 498, 5, 60, 31, 4, 498, 500, 3, 2, 
	2, 2, 499, 492, 3, 2, 2, 2, 499, 495, 3, 2, 2, 2, 500, 503, 3, 2, 2, 2, 
	501, 499, 3, 2, 2, 2, 501, 502, 3, 2, 2, 2, 502, 61, 3, 2, 2, 2, 503, 501, 
	3, 2, 2, 2, 504, 505, 9, 12, 2, 2, 505, 63, 3, 2, 2, 2, 506, 521, 7, 68, 
	2, 2, 507, 521, 7, 69, 2, 2, 508, 521, 7, 70, 2, 2, 509, 510, 7, 70, 2, 
	2, 510, 521, 7, 68, 2, 2, 511, 512, 7, 69, 2, 2, 512, 521, 7, 68, 2, 2, 
	513, 514, 7, 70, 2, 2, 514, 521, 7, 69, 2, 2, 515, 516, 7, 71, 2, 2, 516, 
	521, 7, 68, 2, 2, 517, 518, 7, 70, 2, 2, 518, 519, 7, 68, 2, 2, 519, 521, 
	7, 69, 2, 2, 520, 506, 3, 2, 2, 2, 520, 507, 3, 2, 2, 2, 520, 508, 3, 2, 
	2, 2, 520, 509, 3, 2, 2, 2, 520, 511, 3, 2, 2, 2, 520, 513, 3, 2, 2, 2, 
	520, 515, 3, 2, 2, 2, 520, 517, 3, 2, 2, 2, 521, 65, 3, 2, 2, 2, 522, 523, 
	9, 13, 2, 2, 523, 67, 3, 2, 2, 2, 524, 525, 9, 14, 2, 2, 525, 69, 3, 2, 
	2, 2, 526, 528, 7, 49, 2, 2, 527, 529, 7, 54, 2, 2, 528, 527, 3, 2, 2, 
	2, 528, 529, 3, 2, 2, 2, 529, 531, 3, 2, 2, 2, 530, 532, 7, 25, 2, 2, 531, 
	530, 3, 2, 2, 2, 531, 532, 3, 2, 2, 2, 532, 533, 3, 2, 2, 2, 533, 538, 
	7, 85, 2, 2, 534, 535, 7, 77, 2, 2, 535, 536, 5, 40, 21, 2, 536, 537, 7, 
	78, 2, 2, 537, 539, 3, 2, 2, 2, 538, 534, 3, 2, 2, 2, 538, 539, 3, 2, 2, 
	2, 539, 540, 3, 2, 2, 2, 540, 541, 5, 72, 37, 2, 541, 71, 3, 2, 2, 2, 542, 
	558, 5, 4, 3, 2, 543, 544, 9, 15, 2, 2, 544, 545, 7, 77, 2, 2, 545, 546, 
	5, 76, 39, 2, 546, 554, 7, 78, 2, 2, 547, 548, 7, 79, 2, 2, 548, 549, 7, 
	77, 2, 2, 549, 550, 5, 76, 39, 2, 550, 551, 7, 78, 2, 2, 551, 553, 3, 2, 
	2, 2, 552, 547, 3, 2, 2, 2, 553, 556, 3, 2, 2, 2, 554, 552, 3, 2, 2, 2, 
	554, 555, 3, 2, 2, 2, 555, 558, 3, 2, 2, 2, 556, 554, 3, 2, 2, 2, 557, 
	542, 3, 2, 2, 2, 557, 543, 3, 2, 2, 2, 558, 73, 3, 2, 2, 2, 559, 560, 5, 
	32, 17, 2, 560, 561, 7, 68, 2, 2, 561, 562, 5, 78, 40, 2, 562, 75, 3, 2, 
	2, 2, 563, 568, 5, 78, 40, 2, 564, 565, 7, 79, 2, 2, 565, 567, 5, 78, 40, 
	2, 566, 564, 3, 2, 2, 2, 567, 570, 3, 2, 2, 2, 568, 566, 3, 2, 2, 2, 568, 
	569, 3, 2, 2, 2, 569, 77, 3, 2, 2, 2, 570, 568, 3, 2, 2, 2, 571, 574, 5, 
	56, 29, 2, 572, 574, 7, 53, 2, 2, 573, 571, 3, 2, 2, 2, 573, 572, 3, 2, 
	2, 2, 574, 79, 3, 2, 2, 2, 575, 577, 7, 48, 2, 2, 576, 578, 7, 54, 2, 2, 
	577, 576, 3, 2, 2, 2, 577, 578, 3, 2, 2, 2, 578, 579, 3, 2, 2, 2, 579, 
	580, 7, 20, 2, 2, 580, 583, 7, 85, 2, 2, 581, 582, 7, 47, 2, 2, 582, 584, 
	5, 56, 29, 2, 583, 581, 3, 2, 2, 2, 583, 584, 3, 2, 2, 2, 584, 586, 3, 
	2, 2, 2, 585, 587, 5, 6, 4, 2, 586, 585, 3, 2, 2, 2, 586, 587, 3, 2, 2, 
	2, 587, 589, 3, 2, 2, 2, 588, 590, 5, 30, 16, 2, 589, 588, 3, 2, 2, 2, 
	589, 590, 3, 2, 2, 2, 590, 81, 3, 2, 2, 2, 591, 593, 7, 51, 2, 2, 592, 
	594, 7, 54, 2, 2, 593, 592, 3, 2, 2, 2, 593, 594, 3, 2, 2, 2, 594, 595, 
	3, 2, 2, 2, 595, 600, 7, 85, 2, 2, 596, 598, 7, 9, 2, 2, 597, 596, 3, 2, 
	2, 2, 597, 598, 3, 2, 2, 2, 598, 599, 3, 2, 2, 2, 599, 601, 7, 85, 2, 2, 
	600, 597, 3, 2, 2, 2, 600, 601, 3, 2, 2, 2, 601, 602, 3, 2, 2, 2, 602, 
	603, 7, 50, 2, 2, 603, 608, 5, 74, 38, 2, 604, 605, 7, 79, 2, 2, 605, 607, 
	5, 74, 38, 2, 606, 604, 3, 2, 2, 2, 607, 610, 3, 2, 2, 2, 608, 606, 3, 
	2, 2, 2, 608, 609, 3, 2, 2, 2, 609, 613, 3, 2, 2, 2, 610, 608, 3, 2, 2, 
	2, 611, 612, 7, 47, 2, 2, 612, 614, 5, 56, 29, 2, 613, 611, 3, 2, 2, 2, 
	613, 614, 3, 2, 2, 2, 614, 616, 3, 2, 2, 2, 615, 617, 5, 6, 4, 2, 616, 
	615, 3, 2, 2, 2, 616, 617, 3, 2, 2, 2, 617, 619, 3, 2, 2, 2, 618, 620, 
	5, 30, 16, 2, 619, 618, 3, 2, 2, 2, 619, 620, 3, 2, 2, 2, 620, 83, 3, 2, 
	2, 2, 90, 88, 92, 101, 106, 112, 120, 125, 132, 137, 140, 147, 150, 158, 
	161, 167, 171, 182, 184, 194, 198, 202, 205, 208, 212, 218, 223, 226, 230, 
	233, 237, 240, 242, 248, 257, 262, 264, 268, 273, 279, 285, 291, 302, 305, 
	312, 320, 327, 330, 337, 341, 349, 353, 357, 369, 378, 384, 392, 397, 405, 
	417, 426, 432, 438, 444, 458, 460, 476, 490, 499, 501, 520, 528, 531, 538, 
	554, 557, 568, 573, 577, 583, 586, 589, 593, 597, 600, 608, 613, 616, 619,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", 
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", 
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", 
	"", "", "", "", "'*'", "'/'", "'%'", "'+'", "'--'", "'-'", "", "", "'='", 
	"'>'", "'<'", "'!'", "'~'", "'|'", "'&'", "'^'", "'.'", "'('", "')'", "','",
}
var symbolicNames = []string{
	"", "SPACE", "SPEC_MYSQL_COMMENT", "COMMENT_INPUT", "LINE_COMMENT", "ALL", 
	"AND", "AS", "ASC", "BETWEEN", "BY", "CASE", "CROSS", "DESC", "DISTINCT", 
	"ELSE", "EXISTS", "FALSE", "FROM", "GROUP", "HAVING", "IN", "INNER", "INTO", 
	"IS", "JOIN", "LEFT", "LIKE", "LIMIT", "NOT", "NULL_LITERAL", "ON", "OR", 
	"ORDER", "OUTER", "REGEXP", "RIGHT", "RLIKE", "SELECT", "THEN", "TRUE", 
	"UNION", "USING", "VALUES", "WHEN", "WHERE", "DELETE", "INSERT", "SET", 
	"UPDATE", "VALUE", "DEFAULT", "IGNORE", "ANY", "END", "OFFSET", "SOME", 
	"UNKNOWN", "STAR", "DIVIDE", "MODULE", "PLUS", "MINUSMINUS", "MINUS", "DIV", 
	"MOD", "EQUAL_SYMBOL", "GREATER_SYMBOL", "LESS_SYMBOL", "EXCLAMATION_SYMBOL", 
	"BIT_NOT_OP", "BIT_OR_OP", "BIT_AND_OP", "BIT_XOR_OP", "DOT", "LR_BRACKET", 
	"RR_BRACKET", "COMMA", "STRING_LITERAL", "DECIMAL_LITERAL", "REAL_LITERAL", 
	"NULL_SPEC_LITERAL", "DOT_ID", "ID", "ERROR_RECONGNIGION",
}

var ruleNames = []string{
	"dmlStatement", "selectStatement", "orderByClause", "orderByExpression", 
	"tableSource", "tableSources", "tableSourceItem", "joinPart", "queryExpression", 
	"querySpecification", "selectElements", "selectElement", "fromClause", 
	"groupByItem", "limitClause", "fullColumnName", "booleanLiteral", "nullNotnull", 
	"constant", "uidList", "expressions", "functionCall", "specificFunction", 
	"caseFuncAlternative", "functionArgs", "allFunctionArg", "functionArg", 
	"expression", "predicate", "expressionAtom", "unaryOperator", "comparisonOperator", 
	"logicalOperator", "mathOperator", "insertStatement", "insertStatementValue", 
	"updatedElement", "expressionsWithDefaults", "expressionOrDefault", "deleteStatement", 
	"updateStatement",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type MySqlParser struct {
	*antlr.BaseParser
}

func NewMySqlParser(input antlr.TokenStream) *MySqlParser {
	this := new(MySqlParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "MySqlParser.g4"

	return this
}

// MySqlParser tokens.
const (
	MySqlParserEOF = antlr.TokenEOF
	MySqlParserSPACE = 1
	MySqlParserSPEC_MYSQL_COMMENT = 2
	MySqlParserCOMMENT_INPUT = 3
	MySqlParserLINE_COMMENT = 4
	MySqlParserALL = 5
	MySqlParserAND = 6
	MySqlParserAS = 7
	MySqlParserASC = 8
	MySqlParserBETWEEN = 9
	MySqlParserBY = 10
	MySqlParserCASE = 11
	MySqlParserCROSS = 12
	MySqlParserDESC = 13
	MySqlParserDISTINCT = 14
	MySqlParserELSE = 15
	MySqlParserEXISTS = 16
	MySqlParserFALSE = 17
	MySqlParserFROM = 18
	MySqlParserGROUP = 19
	MySqlParserHAVING = 20
	MySqlParserIN = 21
	MySqlParserINNER = 22
	MySqlParserINTO = 23
	MySqlParserIS = 24
	MySqlParserJOIN = 25
	MySqlParserLEFT = 26
	MySqlParserLIKE = 27
	MySqlParserLIMIT = 28
	MySqlParserNOT = 29
	MySqlParserNULL_LITERAL = 30
	MySqlParserON = 31
	MySqlParserOR = 32
	MySqlParserORDER = 33
	MySqlParserOUTER = 34
	MySqlParserREGEXP = 35
	MySqlParserRIGHT = 36
	MySqlParserRLIKE = 37
	MySqlParserSELECT = 38
	MySqlParserTHEN = 39
	MySqlParserTRUE = 40
	MySqlParserUNION = 41
	MySqlParserUSING = 42
	MySqlParserVALUES = 43
	MySqlParserWHEN = 44
	MySqlParserWHERE = 45
	MySqlParserDELETE = 46
	MySqlParserINSERT = 47
	MySqlParserSET = 48
	MySqlParserUPDATE = 49
	MySqlParserVALUE = 50
	MySqlParserDEFAULT = 51
	MySqlParserIGNORE = 52
	MySqlParserANY = 53
	MySqlParserEND = 54
	MySqlParserOFFSET = 55
	MySqlParserSOME = 56
	MySqlParserUNKNOWN = 57
	MySqlParserSTAR = 58
	MySqlParserDIVIDE = 59
	MySqlParserMODULE = 60
	MySqlParserPLUS = 61
	MySqlParserMINUSMINUS = 62
	MySqlParserMINUS = 63
	MySqlParserDIV = 64
	MySqlParserMOD = 65
	MySqlParserEQUAL_SYMBOL = 66
	MySqlParserGREATER_SYMBOL = 67
	MySqlParserLESS_SYMBOL = 68
	MySqlParserEXCLAMATION_SYMBOL = 69
	MySqlParserBIT_NOT_OP = 70
	MySqlParserBIT_OR_OP = 71
	MySqlParserBIT_AND_OP = 72
	MySqlParserBIT_XOR_OP = 73
	MySqlParserDOT = 74
	MySqlParserLR_BRACKET = 75
	MySqlParserRR_BRACKET = 76
	MySqlParserCOMMA = 77
	MySqlParserSTRING_LITERAL = 78
	MySqlParserDECIMAL_LITERAL = 79
	MySqlParserREAL_LITERAL = 80
	MySqlParserNULL_SPEC_LITERAL = 81
	MySqlParserDOT_ID = 82
	MySqlParserID = 83
	MySqlParserERROR_RECONGNIGION = 84
)

// MySqlParser rules.
const (
	MySqlParserRULE_dmlStatement = 0
	MySqlParserRULE_selectStatement = 1
	MySqlParserRULE_orderByClause = 2
	MySqlParserRULE_orderByExpression = 3
	MySqlParserRULE_tableSource = 4
	MySqlParserRULE_tableSources = 5
	MySqlParserRULE_tableSourceItem = 6
	MySqlParserRULE_joinPart = 7
	MySqlParserRULE_queryExpression = 8
	MySqlParserRULE_querySpecification = 9
	MySqlParserRULE_selectElements = 10
	MySqlParserRULE_selectElement = 11
	MySqlParserRULE_fromClause = 12
	MySqlParserRULE_groupByItem = 13
	MySqlParserRULE_limitClause = 14
	MySqlParserRULE_fullColumnName = 15
	MySqlParserRULE_booleanLiteral = 16
	MySqlParserRULE_nullNotnull = 17
	MySqlParserRULE_constant = 18
	MySqlParserRULE_uidList = 19
	MySqlParserRULE_expressions = 20
	MySqlParserRULE_functionCall = 21
	MySqlParserRULE_specificFunction = 22
	MySqlParserRULE_caseFuncAlternative = 23
	MySqlParserRULE_functionArgs = 24
	MySqlParserRULE_allFunctionArg = 25
	MySqlParserRULE_functionArg = 26
	MySqlParserRULE_expression = 27
	MySqlParserRULE_predicate = 28
	MySqlParserRULE_expressionAtom = 29
	MySqlParserRULE_unaryOperator = 30
	MySqlParserRULE_comparisonOperator = 31
	MySqlParserRULE_logicalOperator = 32
	MySqlParserRULE_mathOperator = 33
	MySqlParserRULE_insertStatement = 34
	MySqlParserRULE_insertStatementValue = 35
	MySqlParserRULE_updatedElement = 36
	MySqlParserRULE_expressionsWithDefaults = 37
	MySqlParserRULE_expressionOrDefault = 38
	MySqlParserRULE_deleteStatement = 39
	MySqlParserRULE_updateStatement = 40
)

// IDmlStatementContext is an interface to support dynamic dispatch.
type IDmlStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDmlStatementContext differentiates from other interfaces.
	IsDmlStatementContext()
}

type DmlStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDmlStatementContext() *DmlStatementContext {
	var p = new(DmlStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySqlParserRULE_dmlStatement
	return p
}

func (*DmlStatementContext) IsDmlStatementContext() {}

func NewDmlStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DmlStatementContext {
	var p = new(DmlStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySqlParserRULE_dmlStatement

	return p
}

func (s *DmlStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *DmlStatementContext) QuerySpecification() IQuerySpecificationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuerySpecificationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuerySpecificationContext)
}

func (s *DmlStatementContext) InsertStatement() IInsertStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInsertStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInsertStatementContext)
}

func (s *DmlStatementContext) UpdateStatement() IUpdateStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUpdateStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUpdateStatementContext)
}

func (s *DmlStatementContext) DeleteStatement() IDeleteStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeleteStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeleteStatementContext)
}

func (s *DmlStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DmlStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *DmlStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterDmlStatement(s)
	}
}

func (s *DmlStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitDmlStatement(s)
	}
}

func (s *DmlStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitDmlStatement(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MySqlParser) DmlStatement() (localctx IDmlStatementContext) {
	localctx = NewDmlStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MySqlParserRULE_dmlStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(86)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MySqlParserSELECT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(82)
			p.QuerySpecification()
		}


	case MySqlParserINSERT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(83)
			p.InsertStatement()
		}


	case MySqlParserUPDATE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(84)
			p.UpdateStatement()
		}


	case MySqlParserDELETE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(85)
			p.DeleteStatement()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// ISelectStatementContext is an interface to support dynamic dispatch.
type ISelectStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectStatementContext differentiates from other interfaces.
	IsSelectStatementContext()
}

type SelectStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectStatementContext() *SelectStatementContext {
	var p = new(SelectStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySqlParserRULE_selectStatement
	return p
}

func (*SelectStatementContext) IsSelectStatementContext() {}

func NewSelectStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectStatementContext {
	var p = new(SelectStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySqlParserRULE_selectStatement

	return p
}

func (s *SelectStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectStatementContext) CopyFrom(ctx *SelectStatementContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *SelectStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




type SimpleSelectContext struct {
	*SelectStatementContext
}

func NewSimpleSelectContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SimpleSelectContext {
	var p = new(SimpleSelectContext)

	p.SelectStatementContext = NewEmptySelectStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SelectStatementContext))

	return p
}

func (s *SimpleSelectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleSelectContext) QuerySpecification() IQuerySpecificationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuerySpecificationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuerySpecificationContext)
}


func (s *SimpleSelectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterSimpleSelect(s)
	}
}

func (s *SimpleSelectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitSimpleSelect(s)
	}
}

func (s *SimpleSelectContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitSimpleSelect(s)

	default:
		return t.VisitChildren(s)
	}
}


type ParenthesisSelectContext struct {
	*SelectStatementContext
}

func NewParenthesisSelectContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenthesisSelectContext {
	var p = new(ParenthesisSelectContext)

	p.SelectStatementContext = NewEmptySelectStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SelectStatementContext))

	return p
}

func (s *ParenthesisSelectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenthesisSelectContext) QueryExpression() IQueryExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQueryExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQueryExpressionContext)
}


func (s *ParenthesisSelectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterParenthesisSelect(s)
	}
}

func (s *ParenthesisSelectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitParenthesisSelect(s)
	}
}

func (s *ParenthesisSelectContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitParenthesisSelect(s)

	default:
		return t.VisitChildren(s)
	}
}



func (p *MySqlParser) SelectStatement() (localctx ISelectStatementContext) {
	localctx = NewSelectStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MySqlParserRULE_selectStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(90)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MySqlParserSELECT:
		localctx = NewSimpleSelectContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(88)
			p.QuerySpecification()
		}


	case MySqlParserLR_BRACKET:
		localctx = NewParenthesisSelectContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(89)
			p.QueryExpression()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IOrderByClauseContext is an interface to support dynamic dispatch.
type IOrderByClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOrderByClauseContext differentiates from other interfaces.
	IsOrderByClauseContext()
}

type OrderByClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrderByClauseContext() *OrderByClauseContext {
	var p = new(OrderByClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySqlParserRULE_orderByClause
	return p
}

func (*OrderByClauseContext) IsOrderByClauseContext() {}

func NewOrderByClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrderByClauseContext {
	var p = new(OrderByClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySqlParserRULE_orderByClause

	return p
}

func (s *OrderByClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *OrderByClauseContext) ORDER() antlr.TerminalNode {
	return s.GetToken(MySqlParserORDER, 0)
}

func (s *OrderByClauseContext) BY() antlr.TerminalNode {
	return s.GetToken(MySqlParserBY, 0)
}

func (s *OrderByClauseContext) AllOrderByExpression() []IOrderByExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOrderByExpressionContext)(nil)).Elem())
	var tst = make([]IOrderByExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOrderByExpressionContext)
		}
	}

	return tst
}

func (s *OrderByClauseContext) OrderByExpression(i int) IOrderByExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrderByExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOrderByExpressionContext)
}

func (s *OrderByClauseContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MySqlParserCOMMA)
}

func (s *OrderByClauseContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MySqlParserCOMMA, i)
}

func (s *OrderByClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrderByClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *OrderByClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterOrderByClause(s)
	}
}

func (s *OrderByClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitOrderByClause(s)
	}
}

func (s *OrderByClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitOrderByClause(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MySqlParser) OrderByClause() (localctx IOrderByClauseContext) {
	localctx = NewOrderByClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MySqlParserRULE_orderByClause)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(92)
		p.Match(MySqlParserORDER)
	}
	{
		p.SetState(93)
		p.Match(MySqlParserBY)
	}
	{
		p.SetState(94)
		p.OrderByExpression()
	}
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == MySqlParserCOMMA {
		{
			p.SetState(95)
			p.Match(MySqlParserCOMMA)
		}
		{
			p.SetState(96)
			p.OrderByExpression()
		}


		p.SetState(101)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


// IOrderByExpressionContext is an interface to support dynamic dispatch.
type IOrderByExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOrder returns the order token.
	GetOrder() antlr.Token 


	// SetOrder sets the order token.
	SetOrder(antlr.Token) 


	// IsOrderByExpressionContext differentiates from other interfaces.
	IsOrderByExpressionContext()
}

type OrderByExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	order antlr.Token
}

func NewEmptyOrderByExpressionContext() *OrderByExpressionContext {
	var p = new(OrderByExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySqlParserRULE_orderByExpression
	return p
}

func (*OrderByExpressionContext) IsOrderByExpressionContext() {}

func NewOrderByExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrderByExpressionContext {
	var p = new(OrderByExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySqlParserRULE_orderByExpression

	return p
}

func (s *OrderByExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *OrderByExpressionContext) GetOrder() antlr.Token { return s.order }


func (s *OrderByExpressionContext) SetOrder(v antlr.Token) { s.order = v }


func (s *OrderByExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *OrderByExpressionContext) ASC() antlr.TerminalNode {
	return s.GetToken(MySqlParserASC, 0)
}

func (s *OrderByExpressionContext) DESC() antlr.TerminalNode {
	return s.GetToken(MySqlParserDESC, 0)
}

func (s *OrderByExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrderByExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *OrderByExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterOrderByExpression(s)
	}
}

func (s *OrderByExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitOrderByExpression(s)
	}
}

func (s *OrderByExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitOrderByExpression(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MySqlParser) OrderByExpression() (localctx IOrderByExpressionContext) {
	localctx = NewOrderByExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MySqlParserRULE_orderByExpression)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(102)
		p.expression(0)
	}
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == MySqlParserASC || _la == MySqlParserDESC {
		{
			p.SetState(103)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*OrderByExpressionContext).order = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == MySqlParserASC || _la == MySqlParserDESC) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*OrderByExpressionContext).order = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}



	return localctx
}


// ITableSourceContext is an interface to support dynamic dispatch.
type ITableSourceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTableSourceContext differentiates from other interfaces.
	IsTableSourceContext()
}

type TableSourceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableSourceContext() *TableSourceContext {
	var p = new(TableSourceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySqlParserRULE_tableSource
	return p
}

func (*TableSourceContext) IsTableSourceContext() {}

func NewTableSourceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableSourceContext {
	var p = new(TableSourceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySqlParserRULE_tableSource

	return p
}

func (s *TableSourceContext) GetParser() antlr.Parser { return s.parser }

func (s *TableSourceContext) CopyFrom(ctx *TableSourceContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TableSourceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableSourceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




type TableSourceNestedContext struct {
	*TableSourceContext
}

func NewTableSourceNestedContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TableSourceNestedContext {
	var p = new(TableSourceNestedContext)

	p.TableSourceContext = NewEmptyTableSourceContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TableSourceContext))

	return p
}

func (s *TableSourceNestedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableSourceNestedContext) LR_BRACKET() antlr.TerminalNode {
	return s.GetToken(MySqlParserLR_BRACKET, 0)
}

func (s *TableSourceNestedContext) TableSourceItem() ITableSourceItemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableSourceItemContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableSourceItemContext)
}

func (s *TableSourceNestedContext) RR_BRACKET() antlr.TerminalNode {
	return s.GetToken(MySqlParserRR_BRACKET, 0)
}

func (s *TableSourceNestedContext) AllJoinPart() []IJoinPartContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IJoinPartContext)(nil)).Elem())
	var tst = make([]IJoinPartContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IJoinPartContext)
		}
	}

	return tst
}

func (s *TableSourceNestedContext) JoinPart(i int) IJoinPartContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJoinPartContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IJoinPartContext)
}


func (s *TableSourceNestedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterTableSourceNested(s)
	}
}

func (s *TableSourceNestedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitTableSourceNested(s)
	}
}

func (s *TableSourceNestedContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitTableSourceNested(s)

	default:
		return t.VisitChildren(s)
	}
}


type TableSourceBaseContext struct {
	*TableSourceContext
}

func NewTableSourceBaseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TableSourceBaseContext {
	var p = new(TableSourceBaseContext)

	p.TableSourceContext = NewEmptyTableSourceContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TableSourceContext))

	return p
}

func (s *TableSourceBaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableSourceBaseContext) TableSourceItem() ITableSourceItemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableSourceItemContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableSourceItemContext)
}

func (s *TableSourceBaseContext) AllJoinPart() []IJoinPartContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IJoinPartContext)(nil)).Elem())
	var tst = make([]IJoinPartContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IJoinPartContext)
		}
	}

	return tst
}

func (s *TableSourceBaseContext) JoinPart(i int) IJoinPartContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJoinPartContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IJoinPartContext)
}


func (s *TableSourceBaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterTableSourceBase(s)
	}
}

func (s *TableSourceBaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitTableSourceBase(s)
	}
}

func (s *TableSourceBaseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitTableSourceBase(s)

	default:
		return t.VisitChildren(s)
	}
}



func (p *MySqlParser) TableSource() (localctx ITableSourceContext) {
	localctx = NewTableSourceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MySqlParserRULE_tableSource)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTableSourceBaseContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(106)
			p.TableSourceItem()
		}
		p.SetState(110)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for ((((_la - 12)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 12))) & ((1 << (MySqlParserCROSS - 12)) | (1 << (MySqlParserINNER - 12)) | (1 << (MySqlParserJOIN - 12)) | (1 << (MySqlParserLEFT - 12)) | (1 << (MySqlParserRIGHT - 12)))) != 0) {
			{
				p.SetState(107)
				p.JoinPart()
			}


			p.SetState(112)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}


	case 2:
		localctx = NewTableSourceNestedContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(113)
			p.Match(MySqlParserLR_BRACKET)
		}
		{
			p.SetState(114)
			p.TableSourceItem()
		}
		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for ((((_la - 12)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 12))) & ((1 << (MySqlParserCROSS - 12)) | (1 << (MySqlParserINNER - 12)) | (1 << (MySqlParserJOIN - 12)) | (1 << (MySqlParserLEFT - 12)) | (1 << (MySqlParserRIGHT - 12)))) != 0) {
			{
				p.SetState(115)
				p.JoinPart()
			}


			p.SetState(120)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(121)
			p.Match(MySqlParserRR_BRACKET)
		}

	}


	return localctx
}


// ITableSourcesContext is an interface to support dynamic dispatch.
type ITableSourcesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTableSourcesContext differentiates from other interfaces.
	IsTableSourcesContext()
}

type TableSourcesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableSourcesContext() *TableSourcesContext {
	var p = new(TableSourcesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySqlParserRULE_tableSources
	return p
}

func (*TableSourcesContext) IsTableSourcesContext() {}

func NewTableSourcesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableSourcesContext {
	var p = new(TableSourcesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySqlParserRULE_tableSources

	return p
}

func (s *TableSourcesContext) GetParser() antlr.Parser { return s.parser }

func (s *TableSourcesContext) AllTableSource() []ITableSourceContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITableSourceContext)(nil)).Elem())
	var tst = make([]ITableSourceContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITableSourceContext)
		}
	}

	return tst
}

func (s *TableSourcesContext) TableSource(i int) ITableSourceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableSourceContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITableSourceContext)
}

func (s *TableSourcesContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MySqlParserCOMMA)
}

func (s *TableSourcesContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MySqlParserCOMMA, i)
}

func (s *TableSourcesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableSourcesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TableSourcesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterTableSources(s)
	}
}

func (s *TableSourcesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitTableSources(s)
	}
}

func (s *TableSourcesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitTableSources(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MySqlParser) TableSources() (localctx ITableSourcesContext) {
	localctx = NewTableSourcesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, MySqlParserRULE_tableSources)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(125)
		p.TableSource()
	}
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == MySqlParserCOMMA {
		{
			p.SetState(126)
			p.Match(MySqlParserCOMMA)
		}
		{
			p.SetState(127)
			p.TableSource()
		}


		p.SetState(132)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


// ITableSourceItemContext is an interface to support dynamic dispatch.
type ITableSourceItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTableSourceItemContext differentiates from other interfaces.
	IsTableSourceItemContext()
}

type TableSourceItemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableSourceItemContext() *TableSourceItemContext {
	var p = new(TableSourceItemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySqlParserRULE_tableSourceItem
	return p
}

func (*TableSourceItemContext) IsTableSourceItemContext() {}

func NewTableSourceItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableSourceItemContext {
	var p = new(TableSourceItemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySqlParserRULE_tableSourceItem

	return p
}

func (s *TableSourceItemContext) GetParser() antlr.Parser { return s.parser }

func (s *TableSourceItemContext) CopyFrom(ctx *TableSourceItemContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TableSourceItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableSourceItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




type SubqueryTableItemContext struct {
	*TableSourceItemContext
	parenthesisSubquery ISelectStatementContext 
	alias antlr.Token
}

func NewSubqueryTableItemContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SubqueryTableItemContext {
	var p = new(SubqueryTableItemContext)

	p.TableSourceItemContext = NewEmptyTableSourceItemContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TableSourceItemContext))

	return p
}


func (s *SubqueryTableItemContext) GetAlias() antlr.Token { return s.alias }


func (s *SubqueryTableItemContext) SetAlias(v antlr.Token) { s.alias = v }


func (s *SubqueryTableItemContext) GetParenthesisSubquery() ISelectStatementContext { return s.parenthesisSubquery }


func (s *SubqueryTableItemContext) SetParenthesisSubquery(v ISelectStatementContext) { s.parenthesisSubquery = v }

func (s *SubqueryTableItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubqueryTableItemContext) ID() antlr.TerminalNode {
	return s.GetToken(MySqlParserID, 0)
}

func (s *SubqueryTableItemContext) SelectStatement() ISelectStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectStatementContext)
}

func (s *SubqueryTableItemContext) LR_BRACKET() antlr.TerminalNode {
	return s.GetToken(MySqlParserLR_BRACKET, 0)
}

func (s *SubqueryTableItemContext) RR_BRACKET() antlr.TerminalNode {
	return s.GetToken(MySqlParserRR_BRACKET, 0)
}

func (s *SubqueryTableItemContext) AS() antlr.TerminalNode {
	return s.GetToken(MySqlParserAS, 0)
}


func (s *SubqueryTableItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterSubqueryTableItem(s)
	}
}

func (s *SubqueryTableItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitSubqueryTableItem(s)
	}
}

func (s *SubqueryTableItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitSubqueryTableItem(s)

	default:
		return t.VisitChildren(s)
	}
}


type AtomTableItemContext struct {
	*TableSourceItemContext
	alias antlr.Token
}

func NewAtomTableItemContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AtomTableItemContext {
	var p = new(AtomTableItemContext)

	p.TableSourceItemContext = NewEmptyTableSourceItemContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TableSourceItemContext))

	return p
}


func (s *AtomTableItemContext) GetAlias() antlr.Token { return s.alias }


func (s *AtomTableItemContext) SetAlias(v antlr.Token) { s.alias = v }

func (s *AtomTableItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomTableItemContext) FullColumnName() IFullColumnNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFullColumnNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFullColumnNameContext)
}

func (s *AtomTableItemContext) ID() antlr.TerminalNode {
	return s.GetToken(MySqlParserID, 0)
}

func (s *AtomTableItemContext) AS() antlr.TerminalNode {
	return s.GetToken(MySqlParserAS, 0)
}


func (s *AtomTableItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterAtomTableItem(s)
	}
}

func (s *AtomTableItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitAtomTableItem(s)
	}
}

func (s *AtomTableItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitAtomTableItem(s)

	default:
		return t.VisitChildren(s)
	}
}


type TableSourcesItemContext struct {
	*TableSourceItemContext
}

func NewTableSourcesItemContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TableSourcesItemContext {
	var p = new(TableSourcesItemContext)

	p.TableSourceItemContext = NewEmptyTableSourceItemContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TableSourceItemContext))

	return p
}

func (s *TableSourcesItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableSourcesItemContext) LR_BRACKET() antlr.TerminalNode {
	return s.GetToken(MySqlParserLR_BRACKET, 0)
}

func (s *TableSourcesItemContext) TableSources() ITableSourcesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableSourcesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableSourcesContext)
}

func (s *TableSourcesItemContext) RR_BRACKET() antlr.TerminalNode {
	return s.GetToken(MySqlParserRR_BRACKET, 0)
}


func (s *TableSourcesItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterTableSourcesItem(s)
	}
}

func (s *TableSourcesItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitTableSourcesItem(s)
	}
}

func (s *TableSourcesItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitTableSourcesItem(s)

	default:
		return t.VisitChildren(s)
	}
}



func (p *MySqlParser) TableSourceItem() (localctx ITableSourceItemContext) {
	localctx = NewTableSourceItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, MySqlParserRULE_tableSourceItem)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAtomTableItemContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(133)
			p.FullColumnName()
		}
		p.SetState(138)
		p.GetErrorHandler().Sync(p)


		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
			p.SetState(135)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)


			if _la == MySqlParserAS {
				{
					p.SetState(134)
					p.Match(MySqlParserAS)
				}

			}
			{
				p.SetState(137)

				var _m = p.Match(MySqlParserID)

				localctx.(*AtomTableItemContext).alias = _m
			}


		}


	case 2:
		localctx = NewSubqueryTableItemContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(145)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(140)
				p.SelectStatement()
			}


		case 2:
			{
				p.SetState(141)
				p.Match(MySqlParserLR_BRACKET)
			}
			{
				p.SetState(142)

				var _x = p.SelectStatement()


				localctx.(*SubqueryTableItemContext).parenthesisSubquery = _x
			}
			{
				p.SetState(143)
				p.Match(MySqlParserRR_BRACKET)
			}

		}
		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		if _la == MySqlParserAS {
			{
				p.SetState(147)
				p.Match(MySqlParserAS)
			}

		}
		{
			p.SetState(150)

			var _m = p.Match(MySqlParserID)

			localctx.(*SubqueryTableItemContext).alias = _m
		}


	case 3:
		localctx = NewTableSourcesItemContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(152)
			p.Match(MySqlParserLR_BRACKET)
		}
		{
			p.SetState(153)
			p.TableSources()
		}
		{
			p.SetState(154)
			p.Match(MySqlParserRR_BRACKET)
		}

	}


	return localctx
}


// IJoinPartContext is an interface to support dynamic dispatch.
type IJoinPartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsJoinPartContext differentiates from other interfaces.
	IsJoinPartContext()
}

type JoinPartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJoinPartContext() *JoinPartContext {
	var p = new(JoinPartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySqlParserRULE_joinPart
	return p
}

func (*JoinPartContext) IsJoinPartContext() {}

func NewJoinPartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JoinPartContext {
	var p = new(JoinPartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySqlParserRULE_joinPart

	return p
}

func (s *JoinPartContext) GetParser() antlr.Parser { return s.parser }

func (s *JoinPartContext) CopyFrom(ctx *JoinPartContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *JoinPartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JoinPartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




type InnerJoinContext struct {
	*JoinPartContext
}

func NewInnerJoinContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InnerJoinContext {
	var p = new(InnerJoinContext)

	p.JoinPartContext = NewEmptyJoinPartContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JoinPartContext))

	return p
}

func (s *InnerJoinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InnerJoinContext) JOIN() antlr.TerminalNode {
	return s.GetToken(MySqlParserJOIN, 0)
}

func (s *InnerJoinContext) TableSourceItem() ITableSourceItemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableSourceItemContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableSourceItemContext)
}

func (s *InnerJoinContext) ON() antlr.TerminalNode {
	return s.GetToken(MySqlParserON, 0)
}

func (s *InnerJoinContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *InnerJoinContext) INNER() antlr.TerminalNode {
	return s.GetToken(MySqlParserINNER, 0)
}

func (s *InnerJoinContext) CROSS() antlr.TerminalNode {
	return s.GetToken(MySqlParserCROSS, 0)
}


func (s *InnerJoinContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterInnerJoin(s)
	}
}

func (s *InnerJoinContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitInnerJoin(s)
	}
}

func (s *InnerJoinContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitInnerJoin(s)

	default:
		return t.VisitChildren(s)
	}
}


type OuterJoinContext struct {
	*JoinPartContext
}

func NewOuterJoinContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OuterJoinContext {
	var p = new(OuterJoinContext)

	p.JoinPartContext = NewEmptyJoinPartContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JoinPartContext))

	return p
}

func (s *OuterJoinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OuterJoinContext) JOIN() antlr.TerminalNode {
	return s.GetToken(MySqlParserJOIN, 0)
}

func (s *OuterJoinContext) TableSourceItem() ITableSourceItemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableSourceItemContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableSourceItemContext)
}

func (s *OuterJoinContext) LEFT() antlr.TerminalNode {
	return s.GetToken(MySqlParserLEFT, 0)
}

func (s *OuterJoinContext) RIGHT() antlr.TerminalNode {
	return s.GetToken(MySqlParserRIGHT, 0)
}

func (s *OuterJoinContext) ON() antlr.TerminalNode {
	return s.GetToken(MySqlParserON, 0)
}

func (s *OuterJoinContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *OuterJoinContext) USING() antlr.TerminalNode {
	return s.GetToken(MySqlParserUSING, 0)
}

func (s *OuterJoinContext) LR_BRACKET() antlr.TerminalNode {
	return s.GetToken(MySqlParserLR_BRACKET, 0)
}

func (s *OuterJoinContext) UidList() IUidListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUidListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUidListContext)
}

func (s *OuterJoinContext) RR_BRACKET() antlr.TerminalNode {
	return s.GetToken(MySqlParserRR_BRACKET, 0)
}

func (s *OuterJoinContext) OUTER() antlr.TerminalNode {
	return s.GetToken(MySqlParserOUTER, 0)
}


func (s *OuterJoinContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterOuterJoin(s)
	}
}

func (s *OuterJoinContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitOuterJoin(s)
	}
}

func (s *OuterJoinContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitOuterJoin(s)

	default:
		return t.VisitChildren(s)
	}
}



func (p *MySqlParser) JoinPart() (localctx IJoinPartContext) {
	localctx = NewJoinPartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, MySqlParserRULE_joinPart)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(182)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MySqlParserCROSS, MySqlParserINNER, MySqlParserJOIN:
		localctx = NewInnerJoinContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(159)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		if _la == MySqlParserCROSS || _la == MySqlParserINNER {
			{
				p.SetState(158)
				_la = p.GetTokenStream().LA(1)

				if !(_la == MySqlParserCROSS || _la == MySqlParserINNER) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		{
			p.SetState(161)
			p.Match(MySqlParserJOIN)
		}
		{
			p.SetState(162)
			p.TableSourceItem()
		}
		p.SetState(165)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		if _la == MySqlParserON {
			{
				p.SetState(163)
				p.Match(MySqlParserON)
			}
			{
				p.SetState(164)
				p.expression(0)
			}

		}


	case MySqlParserLEFT, MySqlParserRIGHT:
		localctx = NewOuterJoinContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(167)
			_la = p.GetTokenStream().LA(1)

			if !(_la == MySqlParserLEFT || _la == MySqlParserRIGHT) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(169)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		if _la == MySqlParserOUTER {
			{
				p.SetState(168)
				p.Match(MySqlParserOUTER)
			}

		}
		{
			p.SetState(171)
			p.Match(MySqlParserJOIN)
		}
		{
			p.SetState(172)
			p.TableSourceItem()
		}
		p.SetState(180)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case MySqlParserON:
			{
				p.SetState(173)
				p.Match(MySqlParserON)
			}
			{
				p.SetState(174)
				p.expression(0)
			}


		case MySqlParserUSING:
			{
				p.SetState(175)
				p.Match(MySqlParserUSING)
			}
			{
				p.SetState(176)
				p.Match(MySqlParserLR_BRACKET)
			}
			{
				p.SetState(177)
				p.UidList()
			}
			{
				p.SetState(178)
				p.Match(MySqlParserRR_BRACKET)
			}



		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IQueryExpressionContext is an interface to support dynamic dispatch.
type IQueryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQueryExpressionContext differentiates from other interfaces.
	IsQueryExpressionContext()
}

type QueryExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueryExpressionContext() *QueryExpressionContext {
	var p = new(QueryExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySqlParserRULE_queryExpression
	return p
}

func (*QueryExpressionContext) IsQueryExpressionContext() {}

func NewQueryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryExpressionContext {
	var p = new(QueryExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySqlParserRULE_queryExpression

	return p
}

func (s *QueryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryExpressionContext) LR_BRACKET() antlr.TerminalNode {
	return s.GetToken(MySqlParserLR_BRACKET, 0)
}

func (s *QueryExpressionContext) QuerySpecification() IQuerySpecificationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuerySpecificationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuerySpecificationContext)
}

func (s *QueryExpressionContext) RR_BRACKET() antlr.TerminalNode {
	return s.GetToken(MySqlParserRR_BRACKET, 0)
}

func (s *QueryExpressionContext) QueryExpression() IQueryExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQueryExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQueryExpressionContext)
}

func (s *QueryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *QueryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterQueryExpression(s)
	}
}

func (s *QueryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitQueryExpression(s)
	}
}

func (s *QueryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitQueryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MySqlParser) QueryExpression() (localctx IQueryExpressionContext) {
	localctx = NewQueryExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, MySqlParserRULE_queryExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(192)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(184)
			p.Match(MySqlParserLR_BRACKET)
		}
		{
			p.SetState(185)
			p.QuerySpecification()
		}
		{
			p.SetState(186)
			p.Match(MySqlParserRR_BRACKET)
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(188)
			p.Match(MySqlParserLR_BRACKET)
		}
		{
			p.SetState(189)
			p.QueryExpression()
		}
		{
			p.SetState(190)
			p.Match(MySqlParserRR_BRACKET)
		}

	}


	return localctx
}


// IQuerySpecificationContext is an interface to support dynamic dispatch.
type IQuerySpecificationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuerySpecificationContext differentiates from other interfaces.
	IsQuerySpecificationContext()
}

type QuerySpecificationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuerySpecificationContext() *QuerySpecificationContext {
	var p = new(QuerySpecificationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySqlParserRULE_querySpecification
	return p
}

func (*QuerySpecificationContext) IsQuerySpecificationContext() {}

func NewQuerySpecificationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuerySpecificationContext {
	var p = new(QuerySpecificationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySqlParserRULE_querySpecification

	return p
}

func (s *QuerySpecificationContext) GetParser() antlr.Parser { return s.parser }

func (s *QuerySpecificationContext) SELECT() antlr.TerminalNode {
	return s.GetToken(MySqlParserSELECT, 0)
}

func (s *QuerySpecificationContext) SelectElements() ISelectElementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectElementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectElementsContext)
}

func (s *QuerySpecificationContext) DISTINCT() antlr.TerminalNode {
	return s.GetToken(MySqlParserDISTINCT, 0)
}

func (s *QuerySpecificationContext) FromClause() IFromClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFromClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFromClauseContext)
}

func (s *QuerySpecificationContext) OrderByClause() IOrderByClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrderByClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOrderByClauseContext)
}

func (s *QuerySpecificationContext) LimitClause() ILimitClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILimitClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILimitClauseContext)
}

func (s *QuerySpecificationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuerySpecificationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *QuerySpecificationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterQuerySpecification(s)
	}
}

func (s *QuerySpecificationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitQuerySpecification(s)
	}
}

func (s *QuerySpecificationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitQuerySpecification(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MySqlParser) QuerySpecification() (localctx IQuerySpecificationContext) {
	localctx = NewQuerySpecificationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, MySqlParserRULE_querySpecification)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(194)
		p.Match(MySqlParserSELECT)
	}
	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == MySqlParserDISTINCT {
		{
			p.SetState(195)
			p.Match(MySqlParserDISTINCT)
		}

	}
	{
		p.SetState(198)
		p.SelectElements()
	}
	p.SetState(200)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == MySqlParserFROM {
		{
			p.SetState(199)
			p.FromClause()
		}

	}
	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == MySqlParserORDER {
		{
			p.SetState(202)
			p.OrderByClause()
		}

	}
	p.SetState(206)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == MySqlParserLIMIT {
		{
			p.SetState(205)
			p.LimitClause()
		}

	}



	return localctx
}


// ISelectElementsContext is an interface to support dynamic dispatch.
type ISelectElementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetStar returns the star token.
	GetStar() antlr.Token 


	// SetStar sets the star token.
	SetStar(antlr.Token) 


	// IsSelectElementsContext differentiates from other interfaces.
	IsSelectElementsContext()
}

type SelectElementsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	star antlr.Token
}

func NewEmptySelectElementsContext() *SelectElementsContext {
	var p = new(SelectElementsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySqlParserRULE_selectElements
	return p
}

func (*SelectElementsContext) IsSelectElementsContext() {}

func NewSelectElementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectElementsContext {
	var p = new(SelectElementsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySqlParserRULE_selectElements

	return p
}

func (s *SelectElementsContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectElementsContext) GetStar() antlr.Token { return s.star }


func (s *SelectElementsContext) SetStar(v antlr.Token) { s.star = v }


func (s *SelectElementsContext) AllSelectElement() []ISelectElementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISelectElementContext)(nil)).Elem())
	var tst = make([]ISelectElementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISelectElementContext)
		}
	}

	return tst
}

func (s *SelectElementsContext) SelectElement(i int) ISelectElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectElementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISelectElementContext)
}

func (s *SelectElementsContext) STAR() antlr.TerminalNode {
	return s.GetToken(MySqlParserSTAR, 0)
}

func (s *SelectElementsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MySqlParserCOMMA)
}

func (s *SelectElementsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MySqlParserCOMMA, i)
}

func (s *SelectElementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectElementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SelectElementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterSelectElements(s)
	}
}

func (s *SelectElementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitSelectElements(s)
	}
}

func (s *SelectElementsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitSelectElements(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MySqlParser) SelectElements() (localctx ISelectElementsContext) {
	localctx = NewSelectElementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, MySqlParserRULE_selectElements)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(210)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MySqlParserSTAR:
		{
			p.SetState(208)

			var _m = p.Match(MySqlParserSTAR)

			localctx.(*SelectElementsContext).star = _m
		}


	case MySqlParserCASE, MySqlParserEXISTS, MySqlParserFALSE, MySqlParserNOT, MySqlParserNULL_LITERAL, MySqlParserTRUE, MySqlParserPLUS, MySqlParserMINUS, MySqlParserEXCLAMATION_SYMBOL, MySqlParserBIT_NOT_OP, MySqlParserLR_BRACKET, MySqlParserSTRING_LITERAL, MySqlParserDECIMAL_LITERAL, MySqlParserREAL_LITERAL, MySqlParserNULL_SPEC_LITERAL, MySqlParserID:
		{
			p.SetState(209)
			p.SelectElement()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(216)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == MySqlParserCOMMA {
		{
			p.SetState(212)
			p.Match(MySqlParserCOMMA)
		}
		{
			p.SetState(213)
			p.SelectElement()
		}


		p.SetState(218)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


// ISelectElementContext is an interface to support dynamic dispatch.
type ISelectElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectElementContext differentiates from other interfaces.
	IsSelectElementContext()
}

type SelectElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectElementContext() *SelectElementContext {
	var p = new(SelectElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySqlParserRULE_selectElement
	return p
}

func (*SelectElementContext) IsSelectElementContext() {}

func NewSelectElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectElementContext {
	var p = new(SelectElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySqlParserRULE_selectElement

	return p
}

func (s *SelectElementContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectElementContext) CopyFrom(ctx *SelectElementContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *SelectElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




type SelectExpressionElementContext struct {
	*SelectElementContext
}

func NewSelectExpressionElementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SelectExpressionElementContext {
	var p = new(SelectExpressionElementContext)

	p.SelectElementContext = NewEmptySelectElementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SelectElementContext))

	return p
}

func (s *SelectExpressionElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectExpressionElementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SelectExpressionElementContext) ID() antlr.TerminalNode {
	return s.GetToken(MySqlParserID, 0)
}

func (s *SelectExpressionElementContext) AS() antlr.TerminalNode {
	return s.GetToken(MySqlParserAS, 0)
}


func (s *SelectExpressionElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterSelectExpressionElement(s)
	}
}

func (s *SelectExpressionElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitSelectExpressionElement(s)
	}
}

func (s *SelectExpressionElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitSelectExpressionElement(s)

	default:
		return t.VisitChildren(s)
	}
}


type SelectFunctionElementContext struct {
	*SelectElementContext
}

func NewSelectFunctionElementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SelectFunctionElementContext {
	var p = new(SelectFunctionElementContext)

	p.SelectElementContext = NewEmptySelectElementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SelectElementContext))

	return p
}

func (s *SelectFunctionElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectFunctionElementContext) FunctionCall() IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *SelectFunctionElementContext) ID() antlr.TerminalNode {
	return s.GetToken(MySqlParserID, 0)
}

func (s *SelectFunctionElementContext) AS() antlr.TerminalNode {
	return s.GetToken(MySqlParserAS, 0)
}


func (s *SelectFunctionElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterSelectFunctionElement(s)
	}
}

func (s *SelectFunctionElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitSelectFunctionElement(s)
	}
}

func (s *SelectFunctionElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitSelectFunctionElement(s)

	default:
		return t.VisitChildren(s)
	}
}


type SelectColumnElementContext struct {
	*SelectElementContext
}

func NewSelectColumnElementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SelectColumnElementContext {
	var p = new(SelectColumnElementContext)

	p.SelectElementContext = NewEmptySelectElementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SelectElementContext))

	return p
}

func (s *SelectColumnElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectColumnElementContext) FullColumnName() IFullColumnNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFullColumnNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFullColumnNameContext)
}

func (s *SelectColumnElementContext) ID() antlr.TerminalNode {
	return s.GetToken(MySqlParserID, 0)
}

func (s *SelectColumnElementContext) AS() antlr.TerminalNode {
	return s.GetToken(MySqlParserAS, 0)
}


func (s *SelectColumnElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterSelectColumnElement(s)
	}
}

func (s *SelectColumnElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitSelectColumnElement(s)
	}
}

func (s *SelectColumnElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitSelectColumnElement(s)

	default:
		return t.VisitChildren(s)
	}
}



func (p *MySqlParser) SelectElement() (localctx ISelectElementContext) {
	localctx = NewSelectElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, MySqlParserRULE_selectElement)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(240)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSelectColumnElementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(219)
			p.FullColumnName()
		}
		p.SetState(224)
		p.GetErrorHandler().Sync(p)


		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) == 1 {
			p.SetState(221)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)


			if _la == MySqlParserAS {
				{
					p.SetState(220)
					p.Match(MySqlParserAS)
				}

			}
			{
				p.SetState(223)
				p.Match(MySqlParserID)
			}


		}


	case 2:
		localctx = NewSelectFunctionElementContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(226)
			p.FunctionCall()
		}
		p.SetState(231)
		p.GetErrorHandler().Sync(p)


		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) == 1 {
			p.SetState(228)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)


			if _la == MySqlParserAS {
				{
					p.SetState(227)
					p.Match(MySqlParserAS)
				}

			}
			{
				p.SetState(230)
				p.Match(MySqlParserID)
			}


		}


	case 3:
		localctx = NewSelectExpressionElementContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(233)
			p.expression(0)
		}
		p.SetState(238)
		p.GetErrorHandler().Sync(p)


		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) == 1 {
			p.SetState(235)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)


			if _la == MySqlParserAS {
				{
					p.SetState(234)
					p.Match(MySqlParserAS)
				}

			}
			{
				p.SetState(237)
				p.Match(MySqlParserID)
			}


		}

	}


	return localctx
}


// IFromClauseContext is an interface to support dynamic dispatch.
type IFromClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetWhereExpr returns the whereExpr rule contexts.
	GetWhereExpr() IExpressionContext

	// GetHavingExpr returns the havingExpr rule contexts.
	GetHavingExpr() IExpressionContext


	// SetWhereExpr sets the whereExpr rule contexts.
	SetWhereExpr(IExpressionContext)

	// SetHavingExpr sets the havingExpr rule contexts.
	SetHavingExpr(IExpressionContext)


	// IsFromClauseContext differentiates from other interfaces.
	IsFromClauseContext()
}

type FromClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	whereExpr IExpressionContext 
	havingExpr IExpressionContext 
}

func NewEmptyFromClauseContext() *FromClauseContext {
	var p = new(FromClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySqlParserRULE_fromClause
	return p
}

func (*FromClauseContext) IsFromClauseContext() {}

func NewFromClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FromClauseContext {
	var p = new(FromClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySqlParserRULE_fromClause

	return p
}

func (s *FromClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *FromClauseContext) GetWhereExpr() IExpressionContext { return s.whereExpr }

func (s *FromClauseContext) GetHavingExpr() IExpressionContext { return s.havingExpr }


func (s *FromClauseContext) SetWhereExpr(v IExpressionContext) { s.whereExpr = v }

func (s *FromClauseContext) SetHavingExpr(v IExpressionContext) { s.havingExpr = v }


func (s *FromClauseContext) FROM() antlr.TerminalNode {
	return s.GetToken(MySqlParserFROM, 0)
}

func (s *FromClauseContext) TableSources() ITableSourcesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableSourcesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableSourcesContext)
}

func (s *FromClauseContext) WHERE() antlr.TerminalNode {
	return s.GetToken(MySqlParserWHERE, 0)
}

func (s *FromClauseContext) GROUP() antlr.TerminalNode {
	return s.GetToken(MySqlParserGROUP, 0)
}

func (s *FromClauseContext) BY() antlr.TerminalNode {
	return s.GetToken(MySqlParserBY, 0)
}

func (s *FromClauseContext) AllGroupByItem() []IGroupByItemContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IGroupByItemContext)(nil)).Elem())
	var tst = make([]IGroupByItemContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IGroupByItemContext)
		}
	}

	return tst
}

func (s *FromClauseContext) GroupByItem(i int) IGroupByItemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroupByItemContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IGroupByItemContext)
}

func (s *FromClauseContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *FromClauseContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FromClauseContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MySqlParserCOMMA)
}

func (s *FromClauseContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MySqlParserCOMMA, i)
}

func (s *FromClauseContext) HAVING() antlr.TerminalNode {
	return s.GetToken(MySqlParserHAVING, 0)
}

func (s *FromClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FromClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FromClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterFromClause(s)
	}
}

func (s *FromClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitFromClause(s)
	}
}

func (s *FromClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitFromClause(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MySqlParser) FromClause() (localctx IFromClauseContext) {
	localctx = NewFromClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, MySqlParserRULE_fromClause)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(242)
		p.Match(MySqlParserFROM)
	}
	{
		p.SetState(243)
		p.TableSources()
	}
	p.SetState(246)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == MySqlParserWHERE {
		{
			p.SetState(244)
			p.Match(MySqlParserWHERE)
		}
		{
			p.SetState(245)

			var _x = p.expression(0)

			localctx.(*FromClauseContext).whereExpr = _x
		}

	}
	p.SetState(262)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == MySqlParserGROUP {
		{
			p.SetState(248)
			p.Match(MySqlParserGROUP)
		}
		{
			p.SetState(249)
			p.Match(MySqlParserBY)
		}
		{
			p.SetState(250)
			p.GroupByItem()
		}
		p.SetState(255)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for _la == MySqlParserCOMMA {
			{
				p.SetState(251)
				p.Match(MySqlParserCOMMA)
			}
			{
				p.SetState(252)
				p.GroupByItem()
			}


			p.SetState(257)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(260)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		if _la == MySqlParserHAVING {
			{
				p.SetState(258)
				p.Match(MySqlParserHAVING)
			}
			{
				p.SetState(259)

				var _x = p.expression(0)

				localctx.(*FromClauseContext).havingExpr = _x
			}

		}

	}



	return localctx
}


// IGroupByItemContext is an interface to support dynamic dispatch.
type IGroupByItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOrder returns the order token.
	GetOrder() antlr.Token 


	// SetOrder sets the order token.
	SetOrder(antlr.Token) 


	// IsGroupByItemContext differentiates from other interfaces.
	IsGroupByItemContext()
}

type GroupByItemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	order antlr.Token
}

func NewEmptyGroupByItemContext() *GroupByItemContext {
	var p = new(GroupByItemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySqlParserRULE_groupByItem
	return p
}

func (*GroupByItemContext) IsGroupByItemContext() {}

func NewGroupByItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupByItemContext {
	var p = new(GroupByItemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySqlParserRULE_groupByItem

	return p
}

func (s *GroupByItemContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupByItemContext) GetOrder() antlr.Token { return s.order }


func (s *GroupByItemContext) SetOrder(v antlr.Token) { s.order = v }


func (s *GroupByItemContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *GroupByItemContext) ASC() antlr.TerminalNode {
	return s.GetToken(MySqlParserASC, 0)
}

func (s *GroupByItemContext) DESC() antlr.TerminalNode {
	return s.GetToken(MySqlParserDESC, 0)
}

func (s *GroupByItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupByItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *GroupByItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterGroupByItem(s)
	}
}

func (s *GroupByItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitGroupByItem(s)
	}
}

func (s *GroupByItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitGroupByItem(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MySqlParser) GroupByItem() (localctx IGroupByItemContext) {
	localctx = NewGroupByItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, MySqlParserRULE_groupByItem)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(264)
		p.expression(0)
	}
	p.SetState(266)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == MySqlParserASC || _la == MySqlParserDESC {
		{
			p.SetState(265)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*GroupByItemContext).order = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == MySqlParserASC || _la == MySqlParserDESC) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*GroupByItemContext).order = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}



	return localctx
}


// ILimitClauseContext is an interface to support dynamic dispatch.
type ILimitClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOffset returns the offset token.
	GetOffset() antlr.Token 

	// GetLimit returns the limit token.
	GetLimit() antlr.Token 


	// SetOffset sets the offset token.
	SetOffset(antlr.Token) 

	// SetLimit sets the limit token.
	SetLimit(antlr.Token) 


	// IsLimitClauseContext differentiates from other interfaces.
	IsLimitClauseContext()
}

type LimitClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	offset antlr.Token
	limit antlr.Token
}

func NewEmptyLimitClauseContext() *LimitClauseContext {
	var p = new(LimitClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySqlParserRULE_limitClause
	return p
}

func (*LimitClauseContext) IsLimitClauseContext() {}

func NewLimitClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LimitClauseContext {
	var p = new(LimitClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySqlParserRULE_limitClause

	return p
}

func (s *LimitClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *LimitClauseContext) GetOffset() antlr.Token { return s.offset }

func (s *LimitClauseContext) GetLimit() antlr.Token { return s.limit }


func (s *LimitClauseContext) SetOffset(v antlr.Token) { s.offset = v }

func (s *LimitClauseContext) SetLimit(v antlr.Token) { s.limit = v }


func (s *LimitClauseContext) LIMIT() antlr.TerminalNode {
	return s.GetToken(MySqlParserLIMIT, 0)
}

func (s *LimitClauseContext) OFFSET() antlr.TerminalNode {
	return s.GetToken(MySqlParserOFFSET, 0)
}

func (s *LimitClauseContext) AllDECIMAL_LITERAL() []antlr.TerminalNode {
	return s.GetTokens(MySqlParserDECIMAL_LITERAL)
}

func (s *LimitClauseContext) DECIMAL_LITERAL(i int) antlr.TerminalNode {
	return s.GetToken(MySqlParserDECIMAL_LITERAL, i)
}

func (s *LimitClauseContext) COMMA() antlr.TerminalNode {
	return s.GetToken(MySqlParserCOMMA, 0)
}

func (s *LimitClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LimitClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *LimitClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterLimitClause(s)
	}
}

func (s *LimitClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitLimitClause(s)
	}
}

func (s *LimitClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitLimitClause(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MySqlParser) LimitClause() (localctx ILimitClauseContext) {
	localctx = NewLimitClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, MySqlParserRULE_limitClause)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(268)
		p.Match(MySqlParserLIMIT)
	}
	p.SetState(277)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		p.SetState(271)
		p.GetErrorHandler().Sync(p)


		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(269)

				var _m = p.Match(MySqlParserDECIMAL_LITERAL)

				localctx.(*LimitClauseContext).offset = _m
			}
			{
				p.SetState(270)
				p.Match(MySqlParserCOMMA)
			}


		}
		{
			p.SetState(273)

			var _m = p.Match(MySqlParserDECIMAL_LITERAL)

			localctx.(*LimitClauseContext).limit = _m
		}


	case 2:
		{
			p.SetState(274)

			var _m = p.Match(MySqlParserDECIMAL_LITERAL)

			localctx.(*LimitClauseContext).limit = _m
		}
		{
			p.SetState(275)
			p.Match(MySqlParserOFFSET)
		}
		{
			p.SetState(276)

			var _m = p.Match(MySqlParserDECIMAL_LITERAL)

			localctx.(*LimitClauseContext).offset = _m
		}

	}



	return localctx
}


// IFullColumnNameContext is an interface to support dynamic dispatch.
type IFullColumnNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFullColumnNameContext differentiates from other interfaces.
	IsFullColumnNameContext()
}

type FullColumnNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFullColumnNameContext() *FullColumnNameContext {
	var p = new(FullColumnNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySqlParserRULE_fullColumnName
	return p
}

func (*FullColumnNameContext) IsFullColumnNameContext() {}

func NewFullColumnNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FullColumnNameContext {
	var p = new(FullColumnNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySqlParserRULE_fullColumnName

	return p
}

func (s *FullColumnNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FullColumnNameContext) ID() antlr.TerminalNode {
	return s.GetToken(MySqlParserID, 0)
}

func (s *FullColumnNameContext) AllDOT_ID() []antlr.TerminalNode {
	return s.GetTokens(MySqlParserDOT_ID)
}

func (s *FullColumnNameContext) DOT_ID(i int) antlr.TerminalNode {
	return s.GetToken(MySqlParserDOT_ID, i)
}

func (s *FullColumnNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FullColumnNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FullColumnNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterFullColumnName(s)
	}
}

func (s *FullColumnNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitFullColumnName(s)
	}
}

func (s *FullColumnNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitFullColumnName(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MySqlParser) FullColumnName() (localctx IFullColumnNameContext) {
	localctx = NewFullColumnNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, MySqlParserRULE_fullColumnName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(279)
		p.Match(MySqlParserID)
	}
	p.SetState(283)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(280)
				p.Match(MySqlParserDOT_ID)
			}


		}
		p.SetState(285)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext())
	}



	return localctx
}


// IBooleanLiteralContext is an interface to support dynamic dispatch.
type IBooleanLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBooleanLiteralContext differentiates from other interfaces.
	IsBooleanLiteralContext()
}

type BooleanLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBooleanLiteralContext() *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySqlParserRULE_booleanLiteral
	return p
}

func (*BooleanLiteralContext) IsBooleanLiteralContext() {}

func NewBooleanLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySqlParserRULE_booleanLiteral

	return p
}

func (s *BooleanLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanLiteralContext) TRUE() antlr.TerminalNode {
	return s.GetToken(MySqlParserTRUE, 0)
}

func (s *BooleanLiteralContext) FALSE() antlr.TerminalNode {
	return s.GetToken(MySqlParserFALSE, 0)
}

func (s *BooleanLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *BooleanLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterBooleanLiteral(s)
	}
}

func (s *BooleanLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitBooleanLiteral(s)
	}
}

func (s *BooleanLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitBooleanLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MySqlParser) BooleanLiteral() (localctx IBooleanLiteralContext) {
	localctx = NewBooleanLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, MySqlParserRULE_booleanLiteral)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(286)
		_la = p.GetTokenStream().LA(1)

		if !(_la == MySqlParserFALSE || _la == MySqlParserTRUE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


// INullNotnullContext is an interface to support dynamic dispatch.
type INullNotnullContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNullNotnullContext differentiates from other interfaces.
	IsNullNotnullContext()
}

type NullNotnullContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNullNotnullContext() *NullNotnullContext {
	var p = new(NullNotnullContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySqlParserRULE_nullNotnull
	return p
}

func (*NullNotnullContext) IsNullNotnullContext() {}

func NewNullNotnullContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NullNotnullContext {
	var p = new(NullNotnullContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySqlParserRULE_nullNotnull

	return p
}

func (s *NullNotnullContext) GetParser() antlr.Parser { return s.parser }

func (s *NullNotnullContext) NULL_LITERAL() antlr.TerminalNode {
	return s.GetToken(MySqlParserNULL_LITERAL, 0)
}

func (s *NullNotnullContext) NULL_SPEC_LITERAL() antlr.TerminalNode {
	return s.GetToken(MySqlParserNULL_SPEC_LITERAL, 0)
}

func (s *NullNotnullContext) NOT() antlr.TerminalNode {
	return s.GetToken(MySqlParserNOT, 0)
}

func (s *NullNotnullContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullNotnullContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *NullNotnullContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterNullNotnull(s)
	}
}

func (s *NullNotnullContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitNullNotnull(s)
	}
}

func (s *NullNotnullContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitNullNotnull(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MySqlParser) NullNotnull() (localctx INullNotnullContext) {
	localctx = NewNullNotnullContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, MySqlParserRULE_nullNotnull)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(289)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == MySqlParserNOT {
		{
			p.SetState(288)
			p.Match(MySqlParserNOT)
		}

	}
	{
		p.SetState(291)
		_la = p.GetTokenStream().LA(1)

		if !(_la == MySqlParserNULL_LITERAL || _la == MySqlParserNULL_SPEC_LITERAL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNullLiteral returns the nullLiteral token.
	GetNullLiteral() antlr.Token 


	// SetNullLiteral sets the nullLiteral token.
	SetNullLiteral(antlr.Token) 


	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	nullLiteral antlr.Token
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySqlParserRULE_constant
	return p
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySqlParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) GetNullLiteral() antlr.Token { return s.nullLiteral }


func (s *ConstantContext) SetNullLiteral(v antlr.Token) { s.nullLiteral = v }


func (s *ConstantContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(MySqlParserSTRING_LITERAL, 0)
}

func (s *ConstantContext) DECIMAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(MySqlParserDECIMAL_LITERAL, 0)
}

func (s *ConstantContext) MINUS() antlr.TerminalNode {
	return s.GetToken(MySqlParserMINUS, 0)
}

func (s *ConstantContext) BooleanLiteral() IBooleanLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanLiteralContext)
}

func (s *ConstantContext) REAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(MySqlParserREAL_LITERAL, 0)
}

func (s *ConstantContext) NULL_LITERAL() antlr.TerminalNode {
	return s.GetToken(MySqlParserNULL_LITERAL, 0)
}

func (s *ConstantContext) NULL_SPEC_LITERAL() antlr.TerminalNode {
	return s.GetToken(MySqlParserNULL_SPEC_LITERAL, 0)
}

func (s *ConstantContext) NOT() antlr.TerminalNode {
	return s.GetToken(MySqlParserNOT, 0)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitConstant(s)
	}
}

func (s *ConstantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitConstant(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MySqlParser) Constant() (localctx IConstantContext) {
	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, MySqlParserRULE_constant)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(303)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MySqlParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(293)
			p.Match(MySqlParserSTRING_LITERAL)
		}


	case MySqlParserDECIMAL_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(294)
			p.Match(MySqlParserDECIMAL_LITERAL)
		}


	case MySqlParserMINUS:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(295)
			p.Match(MySqlParserMINUS)
		}
		{
			p.SetState(296)
			p.Match(MySqlParserDECIMAL_LITERAL)
		}


	case MySqlParserFALSE, MySqlParserTRUE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(297)
			p.BooleanLiteral()
		}


	case MySqlParserREAL_LITERAL:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(298)
			p.Match(MySqlParserREAL_LITERAL)
		}


	case MySqlParserNOT, MySqlParserNULL_LITERAL, MySqlParserNULL_SPEC_LITERAL:
		p.EnterOuterAlt(localctx, 6)
		p.SetState(300)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		if _la == MySqlParserNOT {
			{
				p.SetState(299)
				p.Match(MySqlParserNOT)
			}

		}
		{
			p.SetState(302)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ConstantContext).nullLiteral = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == MySqlParserNULL_LITERAL || _la == MySqlParserNULL_SPEC_LITERAL) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ConstantContext).nullLiteral = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IUidListContext is an interface to support dynamic dispatch.
type IUidListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUidListContext differentiates from other interfaces.
	IsUidListContext()
}

type UidListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUidListContext() *UidListContext {
	var p = new(UidListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySqlParserRULE_uidList
	return p
}

func (*UidListContext) IsUidListContext() {}

func NewUidListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UidListContext {
	var p = new(UidListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySqlParserRULE_uidList

	return p
}

func (s *UidListContext) GetParser() antlr.Parser { return s.parser }

func (s *UidListContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(MySqlParserID)
}

func (s *UidListContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(MySqlParserID, i)
}

func (s *UidListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MySqlParserCOMMA)
}

func (s *UidListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MySqlParserCOMMA, i)
}

func (s *UidListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UidListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *UidListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterUidList(s)
	}
}

func (s *UidListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitUidList(s)
	}
}

func (s *UidListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitUidList(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MySqlParser) UidList() (localctx IUidListContext) {
	localctx = NewUidListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, MySqlParserRULE_uidList)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(305)
		p.Match(MySqlParserID)
	}
	p.SetState(310)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == MySqlParserCOMMA {
		{
			p.SetState(306)
			p.Match(MySqlParserCOMMA)
		}
		{
			p.SetState(307)
			p.Match(MySqlParserID)
		}


		p.SetState(312)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


// IExpressionsContext is an interface to support dynamic dispatch.
type IExpressionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionsContext differentiates from other interfaces.
	IsExpressionsContext()
}

type ExpressionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionsContext() *ExpressionsContext {
	var p = new(ExpressionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySqlParserRULE_expressions
	return p
}

func (*ExpressionsContext) IsExpressionsContext() {}

func NewExpressionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionsContext {
	var p = new(ExpressionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySqlParserRULE_expressions

	return p
}

func (s *ExpressionsContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionsContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionsContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MySqlParserCOMMA)
}

func (s *ExpressionsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MySqlParserCOMMA, i)
}

func (s *ExpressionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ExpressionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterExpressions(s)
	}
}

func (s *ExpressionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitExpressions(s)
	}
}

func (s *ExpressionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitExpressions(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MySqlParser) Expressions() (localctx IExpressionsContext) {
	localctx = NewExpressionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, MySqlParserRULE_expressions)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(313)
		p.expression(0)
	}
	p.SetState(318)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == MySqlParserCOMMA {
		{
			p.SetState(314)
			p.Match(MySqlParserCOMMA)
		}
		{
			p.SetState(315)
			p.expression(0)
		}


		p.SetState(320)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySqlParserRULE_functionCall
	return p
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySqlParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) CopyFrom(ctx *FunctionCallContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




type SpecificFunctionCallContext struct {
	*FunctionCallContext
}

func NewSpecificFunctionCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SpecificFunctionCallContext {
	var p = new(SpecificFunctionCallContext)

	p.FunctionCallContext = NewEmptyFunctionCallContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FunctionCallContext))

	return p
}

func (s *SpecificFunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpecificFunctionCallContext) SpecificFunction() ISpecificFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISpecificFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISpecificFunctionContext)
}


func (s *SpecificFunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterSpecificFunctionCall(s)
	}
}

func (s *SpecificFunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitSpecificFunctionCall(s)
	}
}

func (s *SpecificFunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitSpecificFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}


type ScalarFunctionCallContext struct {
	*FunctionCallContext
}

func NewScalarFunctionCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ScalarFunctionCallContext {
	var p = new(ScalarFunctionCallContext)

	p.FunctionCallContext = NewEmptyFunctionCallContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FunctionCallContext))

	return p
}

func (s *ScalarFunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScalarFunctionCallContext) ID() antlr.TerminalNode {
	return s.GetToken(MySqlParserID, 0)
}

func (s *ScalarFunctionCallContext) LR_BRACKET() antlr.TerminalNode {
	return s.GetToken(MySqlParserLR_BRACKET, 0)
}

func (s *ScalarFunctionCallContext) RR_BRACKET() antlr.TerminalNode {
	return s.GetToken(MySqlParserRR_BRACKET, 0)
}

func (s *ScalarFunctionCallContext) FunctionArgs() IFunctionArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionArgsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionArgsContext)
}


func (s *ScalarFunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterScalarFunctionCall(s)
	}
}

func (s *ScalarFunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitScalarFunctionCall(s)
	}
}

func (s *ScalarFunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitScalarFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}



func (p *MySqlParser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, MySqlParserRULE_functionCall)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(328)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MySqlParserCASE:
		localctx = NewSpecificFunctionCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(321)
			p.SpecificFunction()
		}


	case MySqlParserID:
		localctx = NewScalarFunctionCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(322)
			p.Match(MySqlParserID)
		}
		{
			p.SetState(323)
			p.Match(MySqlParserLR_BRACKET)
		}
		p.SetState(325)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		if (((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << MySqlParserALL) | (1 << MySqlParserCASE) | (1 << MySqlParserEXISTS) | (1 << MySqlParserFALSE) | (1 << MySqlParserNOT) | (1 << MySqlParserNULL_LITERAL))) != 0) || ((((_la - 40)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 40))) & ((1 << (MySqlParserTRUE - 40)) | (1 << (MySqlParserSTAR - 40)) | (1 << (MySqlParserPLUS - 40)) | (1 << (MySqlParserMINUS - 40)) | (1 << (MySqlParserEXCLAMATION_SYMBOL - 40)) | (1 << (MySqlParserBIT_NOT_OP - 40)))) != 0) || ((((_la - 75)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 75))) & ((1 << (MySqlParserLR_BRACKET - 75)) | (1 << (MySqlParserSTRING_LITERAL - 75)) | (1 << (MySqlParserDECIMAL_LITERAL - 75)) | (1 << (MySqlParserREAL_LITERAL - 75)) | (1 << (MySqlParserNULL_SPEC_LITERAL - 75)) | (1 << (MySqlParserID - 75)))) != 0) {
			{
				p.SetState(324)
				p.FunctionArgs()
			}

		}
		{
			p.SetState(327)
			p.Match(MySqlParserRR_BRACKET)
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// ISpecificFunctionContext is an interface to support dynamic dispatch.
type ISpecificFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSpecificFunctionContext differentiates from other interfaces.
	IsSpecificFunctionContext()
}

type SpecificFunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySpecificFunctionContext() *SpecificFunctionContext {
	var p = new(SpecificFunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySqlParserRULE_specificFunction
	return p
}

func (*SpecificFunctionContext) IsSpecificFunctionContext() {}

func NewSpecificFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpecificFunctionContext {
	var p = new(SpecificFunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySqlParserRULE_specificFunction

	return p
}

func (s *SpecificFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *SpecificFunctionContext) CopyFrom(ctx *SpecificFunctionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *SpecificFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpecificFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




type CaseVarFunctionCallContext struct {
	*SpecificFunctionContext
	elseArg IFunctionArgContext 
}

func NewCaseVarFunctionCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CaseVarFunctionCallContext {
	var p = new(CaseVarFunctionCallContext)

	p.SpecificFunctionContext = NewEmptySpecificFunctionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SpecificFunctionContext))

	return p
}


func (s *CaseVarFunctionCallContext) GetElseArg() IFunctionArgContext { return s.elseArg }


func (s *CaseVarFunctionCallContext) SetElseArg(v IFunctionArgContext) { s.elseArg = v }

func (s *CaseVarFunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseVarFunctionCallContext) CASE() antlr.TerminalNode {
	return s.GetToken(MySqlParserCASE, 0)
}

func (s *CaseVarFunctionCallContext) FullColumnName() IFullColumnNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFullColumnNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFullColumnNameContext)
}

func (s *CaseVarFunctionCallContext) END() antlr.TerminalNode {
	return s.GetToken(MySqlParserEND, 0)
}

func (s *CaseVarFunctionCallContext) AllCaseFuncAlternative() []ICaseFuncAlternativeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICaseFuncAlternativeContext)(nil)).Elem())
	var tst = make([]ICaseFuncAlternativeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICaseFuncAlternativeContext)
		}
	}

	return tst
}

func (s *CaseVarFunctionCallContext) CaseFuncAlternative(i int) ICaseFuncAlternativeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICaseFuncAlternativeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICaseFuncAlternativeContext)
}

func (s *CaseVarFunctionCallContext) ELSE() antlr.TerminalNode {
	return s.GetToken(MySqlParserELSE, 0)
}

func (s *CaseVarFunctionCallContext) FunctionArg() IFunctionArgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionArgContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionArgContext)
}


func (s *CaseVarFunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterCaseVarFunctionCall(s)
	}
}

func (s *CaseVarFunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitCaseVarFunctionCall(s)
	}
}

func (s *CaseVarFunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitCaseVarFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}


type CaseFunctionCallContext struct {
	*SpecificFunctionContext
	elseArg IFunctionArgContext 
}

func NewCaseFunctionCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CaseFunctionCallContext {
	var p = new(CaseFunctionCallContext)

	p.SpecificFunctionContext = NewEmptySpecificFunctionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SpecificFunctionContext))

	return p
}


func (s *CaseFunctionCallContext) GetElseArg() IFunctionArgContext { return s.elseArg }


func (s *CaseFunctionCallContext) SetElseArg(v IFunctionArgContext) { s.elseArg = v }

func (s *CaseFunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseFunctionCallContext) CASE() antlr.TerminalNode {
	return s.GetToken(MySqlParserCASE, 0)
}

func (s *CaseFunctionCallContext) END() antlr.TerminalNode {
	return s.GetToken(MySqlParserEND, 0)
}

func (s *CaseFunctionCallContext) AllCaseFuncAlternative() []ICaseFuncAlternativeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICaseFuncAlternativeContext)(nil)).Elem())
	var tst = make([]ICaseFuncAlternativeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICaseFuncAlternativeContext)
		}
	}

	return tst
}

func (s *CaseFunctionCallContext) CaseFuncAlternative(i int) ICaseFuncAlternativeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICaseFuncAlternativeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICaseFuncAlternativeContext)
}

func (s *CaseFunctionCallContext) ELSE() antlr.TerminalNode {
	return s.GetToken(MySqlParserELSE, 0)
}

func (s *CaseFunctionCallContext) FunctionArg() IFunctionArgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionArgContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionArgContext)
}


func (s *CaseFunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterCaseFunctionCall(s)
	}
}

func (s *CaseFunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitCaseFunctionCall(s)
	}
}

func (s *CaseFunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitCaseFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}



func (p *MySqlParser) SpecificFunction() (localctx ISpecificFunctionContext) {
	localctx = NewSpecificFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, MySqlParserRULE_specificFunction)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(355)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 51, p.GetParserRuleContext()) {
	case 1:
		localctx = NewCaseVarFunctionCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(330)
			p.Match(MySqlParserCASE)
		}
		{
			p.SetState(331)
			p.FullColumnName()
		}
		p.SetState(333)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for ok := true; ok; ok = _la == MySqlParserWHEN {
			{
				p.SetState(332)
				p.CaseFuncAlternative()
			}


			p.SetState(335)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(339)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		if _la == MySqlParserELSE {
			{
				p.SetState(337)
				p.Match(MySqlParserELSE)
			}
			{
				p.SetState(338)

				var _x = p.FunctionArg()


				localctx.(*CaseVarFunctionCallContext).elseArg = _x
			}

		}
		{
			p.SetState(341)
			p.Match(MySqlParserEND)
		}


	case 2:
		localctx = NewCaseFunctionCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(343)
			p.Match(MySqlParserCASE)
		}
		p.SetState(345)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for ok := true; ok; ok = _la == MySqlParserWHEN {
			{
				p.SetState(344)
				p.CaseFuncAlternative()
			}


			p.SetState(347)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(351)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		if _la == MySqlParserELSE {
			{
				p.SetState(349)
				p.Match(MySqlParserELSE)
			}
			{
				p.SetState(350)

				var _x = p.FunctionArg()


				localctx.(*CaseFunctionCallContext).elseArg = _x
			}

		}
		{
			p.SetState(353)
			p.Match(MySqlParserEND)
		}

	}


	return localctx
}


// ICaseFuncAlternativeContext is an interface to support dynamic dispatch.
type ICaseFuncAlternativeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCondition returns the condition rule contexts.
	GetCondition() IFunctionArgContext

	// GetConsequent returns the consequent rule contexts.
	GetConsequent() IFunctionArgContext


	// SetCondition sets the condition rule contexts.
	SetCondition(IFunctionArgContext)

	// SetConsequent sets the consequent rule contexts.
	SetConsequent(IFunctionArgContext)


	// IsCaseFuncAlternativeContext differentiates from other interfaces.
	IsCaseFuncAlternativeContext()
}

type CaseFuncAlternativeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	condition IFunctionArgContext 
	consequent IFunctionArgContext 
}

func NewEmptyCaseFuncAlternativeContext() *CaseFuncAlternativeContext {
	var p = new(CaseFuncAlternativeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySqlParserRULE_caseFuncAlternative
	return p
}

func (*CaseFuncAlternativeContext) IsCaseFuncAlternativeContext() {}

func NewCaseFuncAlternativeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseFuncAlternativeContext {
	var p = new(CaseFuncAlternativeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySqlParserRULE_caseFuncAlternative

	return p
}

func (s *CaseFuncAlternativeContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseFuncAlternativeContext) GetCondition() IFunctionArgContext { return s.condition }

func (s *CaseFuncAlternativeContext) GetConsequent() IFunctionArgContext { return s.consequent }


func (s *CaseFuncAlternativeContext) SetCondition(v IFunctionArgContext) { s.condition = v }

func (s *CaseFuncAlternativeContext) SetConsequent(v IFunctionArgContext) { s.consequent = v }


func (s *CaseFuncAlternativeContext) WHEN() antlr.TerminalNode {
	return s.GetToken(MySqlParserWHEN, 0)
}

func (s *CaseFuncAlternativeContext) THEN() antlr.TerminalNode {
	return s.GetToken(MySqlParserTHEN, 0)
}

func (s *CaseFuncAlternativeContext) AllFunctionArg() []IFunctionArgContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunctionArgContext)(nil)).Elem())
	var tst = make([]IFunctionArgContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunctionArgContext)
		}
	}

	return tst
}

func (s *CaseFuncAlternativeContext) FunctionArg(i int) IFunctionArgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionArgContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunctionArgContext)
}

func (s *CaseFuncAlternativeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseFuncAlternativeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *CaseFuncAlternativeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterCaseFuncAlternative(s)
	}
}

func (s *CaseFuncAlternativeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitCaseFuncAlternative(s)
	}
}

func (s *CaseFuncAlternativeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitCaseFuncAlternative(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MySqlParser) CaseFuncAlternative() (localctx ICaseFuncAlternativeContext) {
	localctx = NewCaseFuncAlternativeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, MySqlParserRULE_caseFuncAlternative)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(357)
		p.Match(MySqlParserWHEN)
	}
	{
		p.SetState(358)

		var _x = p.FunctionArg()


		localctx.(*CaseFuncAlternativeContext).condition = _x
	}
	{
		p.SetState(359)
		p.Match(MySqlParserTHEN)
	}
	{
		p.SetState(360)

		var _x = p.FunctionArg()


		localctx.(*CaseFuncAlternativeContext).consequent = _x
	}



	return localctx
}


// IFunctionArgsContext is an interface to support dynamic dispatch.
type IFunctionArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionArgsContext differentiates from other interfaces.
	IsFunctionArgsContext()
}

type FunctionArgsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionArgsContext() *FunctionArgsContext {
	var p = new(FunctionArgsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySqlParserRULE_functionArgs
	return p
}

func (*FunctionArgsContext) IsFunctionArgsContext() {}

func NewFunctionArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionArgsContext {
	var p = new(FunctionArgsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySqlParserRULE_functionArgs

	return p
}

func (s *FunctionArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionArgsContext) AllAllFunctionArg() []IAllFunctionArgContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAllFunctionArgContext)(nil)).Elem())
	var tst = make([]IAllFunctionArgContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAllFunctionArgContext)
		}
	}

	return tst
}

func (s *FunctionArgsContext) AllFunctionArg(i int) IAllFunctionArgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAllFunctionArgContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAllFunctionArgContext)
}

func (s *FunctionArgsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MySqlParserCOMMA)
}

func (s *FunctionArgsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MySqlParserCOMMA, i)
}

func (s *FunctionArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FunctionArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterFunctionArgs(s)
	}
}

func (s *FunctionArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitFunctionArgs(s)
	}
}

func (s *FunctionArgsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitFunctionArgs(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MySqlParser) FunctionArgs() (localctx IFunctionArgsContext) {
	localctx = NewFunctionArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, MySqlParserRULE_functionArgs)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(362)
		p.AllFunctionArg()
	}

	p.SetState(367)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == MySqlParserCOMMA {
		{
			p.SetState(363)
			p.Match(MySqlParserCOMMA)
		}

		{
			p.SetState(364)
			p.AllFunctionArg()
		}



		p.SetState(369)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


// IAllFunctionArgContext is an interface to support dynamic dispatch.
type IAllFunctionArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAllFunctionArgContext differentiates from other interfaces.
	IsAllFunctionArgContext()
}

type AllFunctionArgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAllFunctionArgContext() *AllFunctionArgContext {
	var p = new(AllFunctionArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySqlParserRULE_allFunctionArg
	return p
}

func (*AllFunctionArgContext) IsAllFunctionArgContext() {}

func NewAllFunctionArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AllFunctionArgContext {
	var p = new(AllFunctionArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySqlParserRULE_allFunctionArg

	return p
}

func (s *AllFunctionArgContext) GetParser() antlr.Parser { return s.parser }

func (s *AllFunctionArgContext) STAR() antlr.TerminalNode {
	return s.GetToken(MySqlParserSTAR, 0)
}

func (s *AllFunctionArgContext) ALL() antlr.TerminalNode {
	return s.GetToken(MySqlParserALL, 0)
}

func (s *AllFunctionArgContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *AllFunctionArgContext) FullColumnName() IFullColumnNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFullColumnNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFullColumnNameContext)
}

func (s *AllFunctionArgContext) FunctionCall() IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *AllFunctionArgContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AllFunctionArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllFunctionArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *AllFunctionArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterAllFunctionArg(s)
	}
}

func (s *AllFunctionArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitAllFunctionArg(s)
	}
}

func (s *AllFunctionArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitAllFunctionArg(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MySqlParser) AllFunctionArg() (localctx IAllFunctionArgContext) {
	localctx = NewAllFunctionArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, MySqlParserRULE_allFunctionArg)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(376)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 53, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(370)
			p.Match(MySqlParserSTAR)
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(371)
			p.Match(MySqlParserALL)
		}


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(372)
			p.Constant()
		}


	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(373)
			p.FullColumnName()
		}


	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(374)
			p.FunctionCall()
		}


	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(375)
			p.expression(0)
		}

	}


	return localctx
}


// IFunctionArgContext is an interface to support dynamic dispatch.
type IFunctionArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionArgContext differentiates from other interfaces.
	IsFunctionArgContext()
}

type FunctionArgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionArgContext() *FunctionArgContext {
	var p = new(FunctionArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySqlParserRULE_functionArg
	return p
}

func (*FunctionArgContext) IsFunctionArgContext() {}

func NewFunctionArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionArgContext {
	var p = new(FunctionArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySqlParserRULE_functionArg

	return p
}

func (s *FunctionArgContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionArgContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *FunctionArgContext) FullColumnName() IFullColumnNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFullColumnNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFullColumnNameContext)
}

func (s *FunctionArgContext) FunctionCall() IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *FunctionArgContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FunctionArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FunctionArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterFunctionArg(s)
	}
}

func (s *FunctionArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitFunctionArg(s)
	}
}

func (s *FunctionArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitFunctionArg(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MySqlParser) FunctionArg() (localctx IFunctionArgContext) {
	localctx = NewFunctionArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, MySqlParserRULE_functionArg)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(382)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 54, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(378)
			p.Constant()
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(379)
			p.FullColumnName()
		}


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(380)
			p.FunctionCall()
		}


	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(381)
			p.expression(0)
		}

	}


	return localctx
}


// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySqlParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySqlParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}





type IsExpressionContext struct {
	*ExpressionContext
	testValue antlr.Token
}

func NewIsExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IsExpressionContext {
	var p = new(IsExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}


func (s *IsExpressionContext) GetTestValue() antlr.Token { return s.testValue }


func (s *IsExpressionContext) SetTestValue(v antlr.Token) { s.testValue = v }

func (s *IsExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsExpressionContext) Predicate() IPredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPredicateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPredicateContext)
}

func (s *IsExpressionContext) IS() antlr.TerminalNode {
	return s.GetToken(MySqlParserIS, 0)
}

func (s *IsExpressionContext) TRUE() antlr.TerminalNode {
	return s.GetToken(MySqlParserTRUE, 0)
}

func (s *IsExpressionContext) FALSE() antlr.TerminalNode {
	return s.GetToken(MySqlParserFALSE, 0)
}

func (s *IsExpressionContext) UNKNOWN() antlr.TerminalNode {
	return s.GetToken(MySqlParserUNKNOWN, 0)
}

func (s *IsExpressionContext) NOT() antlr.TerminalNode {
	return s.GetToken(MySqlParserNOT, 0)
}


func (s *IsExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterIsExpression(s)
	}
}

func (s *IsExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitIsExpression(s)
	}
}

func (s *IsExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitIsExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type NotExpressionContext struct {
	*ExpressionContext
	notOperator antlr.Token
}

func NewNotExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExpressionContext {
	var p = new(NotExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}


func (s *NotExpressionContext) GetNotOperator() antlr.Token { return s.notOperator }


func (s *NotExpressionContext) SetNotOperator(v antlr.Token) { s.notOperator = v }

func (s *NotExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NotExpressionContext) NOT() antlr.TerminalNode {
	return s.GetToken(MySqlParserNOT, 0)
}

func (s *NotExpressionContext) EXCLAMATION_SYMBOL() antlr.TerminalNode {
	return s.GetToken(MySqlParserEXCLAMATION_SYMBOL, 0)
}


func (s *NotExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterNotExpression(s)
	}
}

func (s *NotExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitNotExpression(s)
	}
}

func (s *NotExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitNotExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type LogicalExpressionContext struct {
	*ExpressionContext
}

func NewLogicalExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalExpressionContext {
	var p = new(LogicalExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *LogicalExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *LogicalExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LogicalExpressionContext) LogicalOperator() ILogicalOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogicalOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogicalOperatorContext)
}


func (s *LogicalExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterLogicalExpression(s)
	}
}

func (s *LogicalExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitLogicalExpression(s)
	}
}

func (s *LogicalExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitLogicalExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type PredicateExpressionContext struct {
	*ExpressionContext
}

func NewPredicateExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PredicateExpressionContext {
	var p = new(PredicateExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *PredicateExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PredicateExpressionContext) Predicate() IPredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPredicateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPredicateContext)
}


func (s *PredicateExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterPredicateExpression(s)
	}
}

func (s *PredicateExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitPredicateExpression(s)
	}
}

func (s *PredicateExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitPredicateExpression(s)

	default:
		return t.VisitChildren(s)
	}
}



func (p *MySqlParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *MySqlParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 54
	p.EnterRecursionRule(localctx, 54, MySqlParserRULE_expression, _p)
	var _la int


	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(395)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 56, p.GetParserRuleContext()) {
	case 1:
		localctx = NewNotExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(385)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*NotExpressionContext).notOperator = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == MySqlParserNOT || _la == MySqlParserEXCLAMATION_SYMBOL) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*NotExpressionContext).notOperator = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(386)
			p.expression(4)
		}


	case 2:
		localctx = NewIsExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(387)
			p.predicate(0)
		}
		{
			p.SetState(388)
			p.Match(MySqlParserIS)
		}
		p.SetState(390)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		if _la == MySqlParserNOT {
			{
				p.SetState(389)
				p.Match(MySqlParserNOT)
			}

		}
		{
			p.SetState(392)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*IsExpressionContext).testValue = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == MySqlParserFALSE || _la == MySqlParserTRUE || _la == MySqlParserUNKNOWN) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*IsExpressionContext).testValue = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}


	case 3:
		localctx = NewPredicateExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(394)
			p.predicate(0)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(403)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 57, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewLogicalExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, MySqlParserRULE_expression)
			p.SetState(397)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
			}
			{
				p.SetState(398)
				p.LogicalOperator()
			}
			{
				p.SetState(399)
				p.expression(4)
			}


		}
		p.SetState(405)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 57, p.GetParserRuleContext())
	}



	return localctx
}


// IPredicateContext is an interface to support dynamic dispatch.
type IPredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPredicateContext differentiates from other interfaces.
	IsPredicateContext()
}

type PredicateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPredicateContext() *PredicateContext {
	var p = new(PredicateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySqlParserRULE_predicate
	return p
}

func (*PredicateContext) IsPredicateContext() {}

func NewPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PredicateContext {
	var p = new(PredicateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySqlParserRULE_predicate

	return p
}

func (s *PredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *PredicateContext) CopyFrom(ctx *PredicateContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *PredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}





type ExpressionAtomPredicateContext struct {
	*PredicateContext
}

func NewExpressionAtomPredicateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionAtomPredicateContext {
	var p = new(ExpressionAtomPredicateContext)

	p.PredicateContext = NewEmptyPredicateContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PredicateContext))

	return p
}

func (s *ExpressionAtomPredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionAtomPredicateContext) ExpressionAtom() IExpressionAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionAtomContext)
}


func (s *ExpressionAtomPredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterExpressionAtomPredicate(s)
	}
}

func (s *ExpressionAtomPredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitExpressionAtomPredicate(s)
	}
}

func (s *ExpressionAtomPredicateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitExpressionAtomPredicate(s)

	default:
		return t.VisitChildren(s)
	}
}


type InPredicateContext struct {
	*PredicateContext
}

func NewInPredicateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InPredicateContext {
	var p = new(InPredicateContext)

	p.PredicateContext = NewEmptyPredicateContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PredicateContext))

	return p
}

func (s *InPredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InPredicateContext) Predicate() IPredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPredicateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPredicateContext)
}

func (s *InPredicateContext) IN() antlr.TerminalNode {
	return s.GetToken(MySqlParserIN, 0)
}

func (s *InPredicateContext) LR_BRACKET() antlr.TerminalNode {
	return s.GetToken(MySqlParserLR_BRACKET, 0)
}

func (s *InPredicateContext) RR_BRACKET() antlr.TerminalNode {
	return s.GetToken(MySqlParserRR_BRACKET, 0)
}

func (s *InPredicateContext) SelectStatement() ISelectStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectStatementContext)
}

func (s *InPredicateContext) Expressions() IExpressionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *InPredicateContext) NOT() antlr.TerminalNode {
	return s.GetToken(MySqlParserNOT, 0)
}


func (s *InPredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterInPredicate(s)
	}
}

func (s *InPredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitInPredicate(s)
	}
}

func (s *InPredicateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitInPredicate(s)

	default:
		return t.VisitChildren(s)
	}
}


type SubqueryComparasionPredicateContext struct {
	*PredicateContext
	quantifier antlr.Token
}

func NewSubqueryComparasionPredicateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SubqueryComparasionPredicateContext {
	var p = new(SubqueryComparasionPredicateContext)

	p.PredicateContext = NewEmptyPredicateContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PredicateContext))

	return p
}


func (s *SubqueryComparasionPredicateContext) GetQuantifier() antlr.Token { return s.quantifier }


func (s *SubqueryComparasionPredicateContext) SetQuantifier(v antlr.Token) { s.quantifier = v }

func (s *SubqueryComparasionPredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubqueryComparasionPredicateContext) Predicate() IPredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPredicateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPredicateContext)
}

func (s *SubqueryComparasionPredicateContext) ComparisonOperator() IComparisonOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparisonOperatorContext)
}

func (s *SubqueryComparasionPredicateContext) LR_BRACKET() antlr.TerminalNode {
	return s.GetToken(MySqlParserLR_BRACKET, 0)
}

func (s *SubqueryComparasionPredicateContext) SelectStatement() ISelectStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectStatementContext)
}

func (s *SubqueryComparasionPredicateContext) RR_BRACKET() antlr.TerminalNode {
	return s.GetToken(MySqlParserRR_BRACKET, 0)
}

func (s *SubqueryComparasionPredicateContext) ALL() antlr.TerminalNode {
	return s.GetToken(MySqlParserALL, 0)
}

func (s *SubqueryComparasionPredicateContext) ANY() antlr.TerminalNode {
	return s.GetToken(MySqlParserANY, 0)
}

func (s *SubqueryComparasionPredicateContext) SOME() antlr.TerminalNode {
	return s.GetToken(MySqlParserSOME, 0)
}


func (s *SubqueryComparasionPredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterSubqueryComparasionPredicate(s)
	}
}

func (s *SubqueryComparasionPredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitSubqueryComparasionPredicate(s)
	}
}

func (s *SubqueryComparasionPredicateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitSubqueryComparasionPredicate(s)

	default:
		return t.VisitChildren(s)
	}
}


type BetweenPredicateContext struct {
	*PredicateContext
}

func NewBetweenPredicateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BetweenPredicateContext {
	var p = new(BetweenPredicateContext)

	p.PredicateContext = NewEmptyPredicateContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PredicateContext))

	return p
}

func (s *BetweenPredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BetweenPredicateContext) AllPredicate() []IPredicateContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPredicateContext)(nil)).Elem())
	var tst = make([]IPredicateContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPredicateContext)
		}
	}

	return tst
}

func (s *BetweenPredicateContext) Predicate(i int) IPredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPredicateContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPredicateContext)
}

func (s *BetweenPredicateContext) BETWEEN() antlr.TerminalNode {
	return s.GetToken(MySqlParserBETWEEN, 0)
}

func (s *BetweenPredicateContext) AND() antlr.TerminalNode {
	return s.GetToken(MySqlParserAND, 0)
}

func (s *BetweenPredicateContext) NOT() antlr.TerminalNode {
	return s.GetToken(MySqlParserNOT, 0)
}


func (s *BetweenPredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterBetweenPredicate(s)
	}
}

func (s *BetweenPredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitBetweenPredicate(s)
	}
}

func (s *BetweenPredicateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitBetweenPredicate(s)

	default:
		return t.VisitChildren(s)
	}
}


type BinaryComparasionPredicateContext struct {
	*PredicateContext
	left IPredicateContext 
	right IPredicateContext 
}

func NewBinaryComparasionPredicateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryComparasionPredicateContext {
	var p = new(BinaryComparasionPredicateContext)

	p.PredicateContext = NewEmptyPredicateContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PredicateContext))

	return p
}


func (s *BinaryComparasionPredicateContext) GetLeft() IPredicateContext { return s.left }

func (s *BinaryComparasionPredicateContext) GetRight() IPredicateContext { return s.right }


func (s *BinaryComparasionPredicateContext) SetLeft(v IPredicateContext) { s.left = v }

func (s *BinaryComparasionPredicateContext) SetRight(v IPredicateContext) { s.right = v }

func (s *BinaryComparasionPredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryComparasionPredicateContext) ComparisonOperator() IComparisonOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparisonOperatorContext)
}

func (s *BinaryComparasionPredicateContext) AllPredicate() []IPredicateContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPredicateContext)(nil)).Elem())
	var tst = make([]IPredicateContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPredicateContext)
		}
	}

	return tst
}

func (s *BinaryComparasionPredicateContext) Predicate(i int) IPredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPredicateContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPredicateContext)
}


func (s *BinaryComparasionPredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterBinaryComparasionPredicate(s)
	}
}

func (s *BinaryComparasionPredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitBinaryComparasionPredicate(s)
	}
}

func (s *BinaryComparasionPredicateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitBinaryComparasionPredicate(s)

	default:
		return t.VisitChildren(s)
	}
}


type IsNullPredicateContext struct {
	*PredicateContext
}

func NewIsNullPredicateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IsNullPredicateContext {
	var p = new(IsNullPredicateContext)

	p.PredicateContext = NewEmptyPredicateContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PredicateContext))

	return p
}

func (s *IsNullPredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsNullPredicateContext) Predicate() IPredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPredicateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPredicateContext)
}

func (s *IsNullPredicateContext) IS() antlr.TerminalNode {
	return s.GetToken(MySqlParserIS, 0)
}

func (s *IsNullPredicateContext) NullNotnull() INullNotnullContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INullNotnullContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INullNotnullContext)
}


func (s *IsNullPredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterIsNullPredicate(s)
	}
}

func (s *IsNullPredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitIsNullPredicate(s)
	}
}

func (s *IsNullPredicateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitIsNullPredicate(s)

	default:
		return t.VisitChildren(s)
	}
}


type LikePredicateContext struct {
	*PredicateContext
}

func NewLikePredicateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LikePredicateContext {
	var p = new(LikePredicateContext)

	p.PredicateContext = NewEmptyPredicateContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PredicateContext))

	return p
}

func (s *LikePredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LikePredicateContext) AllPredicate() []IPredicateContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPredicateContext)(nil)).Elem())
	var tst = make([]IPredicateContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPredicateContext)
		}
	}

	return tst
}

func (s *LikePredicateContext) Predicate(i int) IPredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPredicateContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPredicateContext)
}

func (s *LikePredicateContext) LIKE() antlr.TerminalNode {
	return s.GetToken(MySqlParserLIKE, 0)
}

func (s *LikePredicateContext) NOT() antlr.TerminalNode {
	return s.GetToken(MySqlParserNOT, 0)
}


func (s *LikePredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterLikePredicate(s)
	}
}

func (s *LikePredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitLikePredicate(s)
	}
}

func (s *LikePredicateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitLikePredicate(s)

	default:
		return t.VisitChildren(s)
	}
}


type RegexpPredicateContext struct {
	*PredicateContext
	regex antlr.Token
}

func NewRegexpPredicateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RegexpPredicateContext {
	var p = new(RegexpPredicateContext)

	p.PredicateContext = NewEmptyPredicateContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PredicateContext))

	return p
}


func (s *RegexpPredicateContext) GetRegex() antlr.Token { return s.regex }


func (s *RegexpPredicateContext) SetRegex(v antlr.Token) { s.regex = v }

func (s *RegexpPredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegexpPredicateContext) AllPredicate() []IPredicateContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPredicateContext)(nil)).Elem())
	var tst = make([]IPredicateContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPredicateContext)
		}
	}

	return tst
}

func (s *RegexpPredicateContext) Predicate(i int) IPredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPredicateContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPredicateContext)
}

func (s *RegexpPredicateContext) REGEXP() antlr.TerminalNode {
	return s.GetToken(MySqlParserREGEXP, 0)
}

func (s *RegexpPredicateContext) RLIKE() antlr.TerminalNode {
	return s.GetToken(MySqlParserRLIKE, 0)
}

func (s *RegexpPredicateContext) NOT() antlr.TerminalNode {
	return s.GetToken(MySqlParserNOT, 0)
}


func (s *RegexpPredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterRegexpPredicate(s)
	}
}

func (s *RegexpPredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitRegexpPredicate(s)
	}
}

func (s *RegexpPredicateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitRegexpPredicate(s)

	default:
		return t.VisitChildren(s)
	}
}



func (p *MySqlParser) Predicate() (localctx IPredicateContext) {
	return p.predicate(0)
}

func (p *MySqlParser) predicate(_p int) (localctx IPredicateContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewPredicateContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPredicateContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 56
	p.EnterRecursionRule(localctx, 56, MySqlParserRULE_predicate, _p)
	var _la int


	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewExpressionAtomPredicateContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(407)
		p.expressionAtom(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(458)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 64, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(456)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 63, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBinaryComparasionPredicateContext(p, NewPredicateContext(p, _parentctx, _parentState))
				localctx.(*BinaryComparasionPredicateContext).left = _prevctx


				p.PushNewRecursionContext(localctx, _startState, MySqlParserRULE_predicate)
				p.SetState(409)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(410)
					p.ComparisonOperator()
				}
				{
					p.SetState(411)

					var _x = p.predicate(7)

					localctx.(*BinaryComparasionPredicateContext).right = _x
				}


			case 2:
				localctx = NewBetweenPredicateContext(p, NewPredicateContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MySqlParserRULE_predicate)
				p.SetState(413)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				p.SetState(415)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)


				if _la == MySqlParserNOT {
					{
						p.SetState(414)
						p.Match(MySqlParserNOT)
					}

				}
				{
					p.SetState(417)
					p.Match(MySqlParserBETWEEN)
				}
				{
					p.SetState(418)
					p.predicate(0)
				}
				{
					p.SetState(419)
					p.Match(MySqlParserAND)
				}
				{
					p.SetState(420)
					p.predicate(5)
				}


			case 3:
				localctx = NewLikePredicateContext(p, NewPredicateContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MySqlParserRULE_predicate)
				p.SetState(422)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				p.SetState(424)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)


				if _la == MySqlParserNOT {
					{
						p.SetState(423)
						p.Match(MySqlParserNOT)
					}

				}
				{
					p.SetState(426)
					p.Match(MySqlParserLIKE)
				}
				{
					p.SetState(427)
					p.predicate(4)
				}


			case 4:
				localctx = NewRegexpPredicateContext(p, NewPredicateContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MySqlParserRULE_predicate)
				p.SetState(428)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				p.SetState(430)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)


				if _la == MySqlParserNOT {
					{
						p.SetState(429)
						p.Match(MySqlParserNOT)
					}

				}
				{
					p.SetState(432)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*RegexpPredicateContext).regex = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == MySqlParserREGEXP || _la == MySqlParserRLIKE) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*RegexpPredicateContext).regex = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(433)
					p.predicate(3)
				}


			case 5:
				localctx = NewInPredicateContext(p, NewPredicateContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MySqlParserRULE_predicate)
				p.SetState(434)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				p.SetState(436)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)


				if _la == MySqlParserNOT {
					{
						p.SetState(435)
						p.Match(MySqlParserNOT)
					}

				}
				{
					p.SetState(438)
					p.Match(MySqlParserIN)
				}
				{
					p.SetState(439)
					p.Match(MySqlParserLR_BRACKET)
				}
				p.SetState(442)
				p.GetErrorHandler().Sync(p)
				switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 62, p.GetParserRuleContext()) {
				case 1:
					{
						p.SetState(440)
						p.SelectStatement()
					}


				case 2:
					{
						p.SetState(441)
						p.Expressions()
					}

				}
				{
					p.SetState(444)
					p.Match(MySqlParserRR_BRACKET)
				}


			case 6:
				localctx = NewIsNullPredicateContext(p, NewPredicateContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MySqlParserRULE_predicate)
				p.SetState(446)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(447)
					p.Match(MySqlParserIS)
				}
				{
					p.SetState(448)
					p.NullNotnull()
				}


			case 7:
				localctx = NewSubqueryComparasionPredicateContext(p, NewPredicateContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MySqlParserRULE_predicate)
				p.SetState(449)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(450)
					p.ComparisonOperator()
				}
				{
					p.SetState(451)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*SubqueryComparasionPredicateContext).quantifier = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == MySqlParserALL || _la == MySqlParserANY || _la == MySqlParserSOME) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*SubqueryComparasionPredicateContext).quantifier = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(452)
					p.Match(MySqlParserLR_BRACKET)
				}
				{
					p.SetState(453)
					p.SelectStatement()
				}
				{
					p.SetState(454)
					p.Match(MySqlParserRR_BRACKET)
				}

			}

		}
		p.SetState(460)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 64, p.GetParserRuleContext())
	}



	return localctx
}


// IExpressionAtomContext is an interface to support dynamic dispatch.
type IExpressionAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionAtomContext differentiates from other interfaces.
	IsExpressionAtomContext()
}

type ExpressionAtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionAtomContext() *ExpressionAtomContext {
	var p = new(ExpressionAtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySqlParserRULE_expressionAtom
	return p
}

func (*ExpressionAtomContext) IsExpressionAtomContext() {}

func NewExpressionAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionAtomContext {
	var p = new(ExpressionAtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySqlParserRULE_expressionAtom

	return p
}

func (s *ExpressionAtomContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionAtomContext) CopyFrom(ctx *ExpressionAtomContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionAtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}





type UnaryExpressionAtomContext struct {
	*ExpressionAtomContext
}

func NewUnaryExpressionAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryExpressionAtomContext {
	var p = new(UnaryExpressionAtomContext)

	p.ExpressionAtomContext = NewEmptyExpressionAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionAtomContext))

	return p
}

func (s *UnaryExpressionAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExpressionAtomContext) UnaryOperator() IUnaryOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnaryOperatorContext)
}

func (s *UnaryExpressionAtomContext) ExpressionAtom() IExpressionAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionAtomContext)
}


func (s *UnaryExpressionAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterUnaryExpressionAtom(s)
	}
}

func (s *UnaryExpressionAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitUnaryExpressionAtom(s)
	}
}

func (s *UnaryExpressionAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitUnaryExpressionAtom(s)

	default:
		return t.VisitChildren(s)
	}
}


type SubqueryExpressionAtomContext struct {
	*ExpressionAtomContext
}

func NewSubqueryExpressionAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SubqueryExpressionAtomContext {
	var p = new(SubqueryExpressionAtomContext)

	p.ExpressionAtomContext = NewEmptyExpressionAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionAtomContext))

	return p
}

func (s *SubqueryExpressionAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubqueryExpressionAtomContext) LR_BRACKET() antlr.TerminalNode {
	return s.GetToken(MySqlParserLR_BRACKET, 0)
}

func (s *SubqueryExpressionAtomContext) SelectStatement() ISelectStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectStatementContext)
}

func (s *SubqueryExpressionAtomContext) RR_BRACKET() antlr.TerminalNode {
	return s.GetToken(MySqlParserRR_BRACKET, 0)
}


func (s *SubqueryExpressionAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterSubqueryExpressionAtom(s)
	}
}

func (s *SubqueryExpressionAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitSubqueryExpressionAtom(s)
	}
}

func (s *SubqueryExpressionAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitSubqueryExpressionAtom(s)

	default:
		return t.VisitChildren(s)
	}
}


type PriorityMathExpressionAtomContext struct {
	*ExpressionAtomContext
	left IExpressionAtomContext 
	op antlr.Token
	right IExpressionAtomContext 
}

func NewPriorityMathExpressionAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PriorityMathExpressionAtomContext {
	var p = new(PriorityMathExpressionAtomContext)

	p.ExpressionAtomContext = NewEmptyExpressionAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionAtomContext))

	return p
}


func (s *PriorityMathExpressionAtomContext) GetOp() antlr.Token { return s.op }


func (s *PriorityMathExpressionAtomContext) SetOp(v antlr.Token) { s.op = v }


func (s *PriorityMathExpressionAtomContext) GetLeft() IExpressionAtomContext { return s.left }

func (s *PriorityMathExpressionAtomContext) GetRight() IExpressionAtomContext { return s.right }


func (s *PriorityMathExpressionAtomContext) SetLeft(v IExpressionAtomContext) { s.left = v }

func (s *PriorityMathExpressionAtomContext) SetRight(v IExpressionAtomContext) { s.right = v }

func (s *PriorityMathExpressionAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PriorityMathExpressionAtomContext) AllExpressionAtom() []IExpressionAtomContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionAtomContext)(nil)).Elem())
	var tst = make([]IExpressionAtomContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionAtomContext)
		}
	}

	return tst
}

func (s *PriorityMathExpressionAtomContext) ExpressionAtom(i int) IExpressionAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionAtomContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionAtomContext)
}

func (s *PriorityMathExpressionAtomContext) STAR() antlr.TerminalNode {
	return s.GetToken(MySqlParserSTAR, 0)
}

func (s *PriorityMathExpressionAtomContext) DIVIDE() antlr.TerminalNode {
	return s.GetToken(MySqlParserDIVIDE, 0)
}

func (s *PriorityMathExpressionAtomContext) MODULE() antlr.TerminalNode {
	return s.GetToken(MySqlParserMODULE, 0)
}

func (s *PriorityMathExpressionAtomContext) DIV() antlr.TerminalNode {
	return s.GetToken(MySqlParserDIV, 0)
}

func (s *PriorityMathExpressionAtomContext) MOD() antlr.TerminalNode {
	return s.GetToken(MySqlParserMOD, 0)
}


func (s *PriorityMathExpressionAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterPriorityMathExpressionAtom(s)
	}
}

func (s *PriorityMathExpressionAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitPriorityMathExpressionAtom(s)
	}
}

func (s *PriorityMathExpressionAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitPriorityMathExpressionAtom(s)

	default:
		return t.VisitChildren(s)
	}
}


type ConstantExpressionAtomContext struct {
	*ExpressionAtomContext
}

func NewConstantExpressionAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConstantExpressionAtomContext {
	var p = new(ConstantExpressionAtomContext)

	p.ExpressionAtomContext = NewEmptyExpressionAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionAtomContext))

	return p
}

func (s *ConstantExpressionAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantExpressionAtomContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}


func (s *ConstantExpressionAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterConstantExpressionAtom(s)
	}
}

func (s *ConstantExpressionAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitConstantExpressionAtom(s)
	}
}

func (s *ConstantExpressionAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitConstantExpressionAtom(s)

	default:
		return t.VisitChildren(s)
	}
}


type FunctionCallExpressionAtomContext struct {
	*ExpressionAtomContext
}

func NewFunctionCallExpressionAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionCallExpressionAtomContext {
	var p = new(FunctionCallExpressionAtomContext)

	p.ExpressionAtomContext = NewEmptyExpressionAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionAtomContext))

	return p
}

func (s *FunctionCallExpressionAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallExpressionAtomContext) FunctionCall() IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}


func (s *FunctionCallExpressionAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterFunctionCallExpressionAtom(s)
	}
}

func (s *FunctionCallExpressionAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitFunctionCallExpressionAtom(s)
	}
}

func (s *FunctionCallExpressionAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitFunctionCallExpressionAtom(s)

	default:
		return t.VisitChildren(s)
	}
}


type FullColumnNameExpressionAtomContext struct {
	*ExpressionAtomContext
}

func NewFullColumnNameExpressionAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FullColumnNameExpressionAtomContext {
	var p = new(FullColumnNameExpressionAtomContext)

	p.ExpressionAtomContext = NewEmptyExpressionAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionAtomContext))

	return p
}

func (s *FullColumnNameExpressionAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FullColumnNameExpressionAtomContext) FullColumnName() IFullColumnNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFullColumnNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFullColumnNameContext)
}


func (s *FullColumnNameExpressionAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterFullColumnNameExpressionAtom(s)
	}
}

func (s *FullColumnNameExpressionAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitFullColumnNameExpressionAtom(s)
	}
}

func (s *FullColumnNameExpressionAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitFullColumnNameExpressionAtom(s)

	default:
		return t.VisitChildren(s)
	}
}


type NestedExpressionAtomContext struct {
	*ExpressionAtomContext
}

func NewNestedExpressionAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NestedExpressionAtomContext {
	var p = new(NestedExpressionAtomContext)

	p.ExpressionAtomContext = NewEmptyExpressionAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionAtomContext))

	return p
}

func (s *NestedExpressionAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NestedExpressionAtomContext) LR_BRACKET() antlr.TerminalNode {
	return s.GetToken(MySqlParserLR_BRACKET, 0)
}

func (s *NestedExpressionAtomContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *NestedExpressionAtomContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NestedExpressionAtomContext) RR_BRACKET() antlr.TerminalNode {
	return s.GetToken(MySqlParserRR_BRACKET, 0)
}

func (s *NestedExpressionAtomContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MySqlParserCOMMA)
}

func (s *NestedExpressionAtomContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MySqlParserCOMMA, i)
}


func (s *NestedExpressionAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterNestedExpressionAtom(s)
	}
}

func (s *NestedExpressionAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitNestedExpressionAtom(s)
	}
}

func (s *NestedExpressionAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitNestedExpressionAtom(s)

	default:
		return t.VisitChildren(s)
	}
}


type MathExpressionAtomContext struct {
	*ExpressionAtomContext
	left IExpressionAtomContext 
	right IExpressionAtomContext 
}

func NewMathExpressionAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MathExpressionAtomContext {
	var p = new(MathExpressionAtomContext)

	p.ExpressionAtomContext = NewEmptyExpressionAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionAtomContext))

	return p
}


func (s *MathExpressionAtomContext) GetLeft() IExpressionAtomContext { return s.left }

func (s *MathExpressionAtomContext) GetRight() IExpressionAtomContext { return s.right }


func (s *MathExpressionAtomContext) SetLeft(v IExpressionAtomContext) { s.left = v }

func (s *MathExpressionAtomContext) SetRight(v IExpressionAtomContext) { s.right = v }

func (s *MathExpressionAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MathExpressionAtomContext) MathOperator() IMathOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathOperatorContext)
}

func (s *MathExpressionAtomContext) AllExpressionAtom() []IExpressionAtomContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionAtomContext)(nil)).Elem())
	var tst = make([]IExpressionAtomContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionAtomContext)
		}
	}

	return tst
}

func (s *MathExpressionAtomContext) ExpressionAtom(i int) IExpressionAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionAtomContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionAtomContext)
}


func (s *MathExpressionAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterMathExpressionAtom(s)
	}
}

func (s *MathExpressionAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitMathExpressionAtom(s)
	}
}

func (s *MathExpressionAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitMathExpressionAtom(s)

	default:
		return t.VisitChildren(s)
	}
}


type ExistsExpressionAtomContext struct {
	*ExpressionAtomContext
}

func NewExistsExpressionAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExistsExpressionAtomContext {
	var p = new(ExistsExpressionAtomContext)

	p.ExpressionAtomContext = NewEmptyExpressionAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionAtomContext))

	return p
}

func (s *ExistsExpressionAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExistsExpressionAtomContext) EXISTS() antlr.TerminalNode {
	return s.GetToken(MySqlParserEXISTS, 0)
}

func (s *ExistsExpressionAtomContext) LR_BRACKET() antlr.TerminalNode {
	return s.GetToken(MySqlParserLR_BRACKET, 0)
}

func (s *ExistsExpressionAtomContext) SelectStatement() ISelectStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectStatementContext)
}

func (s *ExistsExpressionAtomContext) RR_BRACKET() antlr.TerminalNode {
	return s.GetToken(MySqlParserRR_BRACKET, 0)
}


func (s *ExistsExpressionAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterExistsExpressionAtom(s)
	}
}

func (s *ExistsExpressionAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitExistsExpressionAtom(s)
	}
}

func (s *ExistsExpressionAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitExistsExpressionAtom(s)

	default:
		return t.VisitChildren(s)
	}
}



func (p *MySqlParser) ExpressionAtom() (localctx IExpressionAtomContext) {
	return p.expressionAtom(0)
}

func (p *MySqlParser) expressionAtom(_p int) (localctx IExpressionAtomContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionAtomContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionAtomContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 58
	p.EnterRecursionRule(localctx, 58, MySqlParserRULE_expressionAtom, _p)
	var _la int


	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(488)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 66, p.GetParserRuleContext()) {
	case 1:
		localctx = NewConstantExpressionAtomContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(462)
			p.Constant()
		}


	case 2:
		localctx = NewFullColumnNameExpressionAtomContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(463)
			p.FullColumnName()
		}


	case 3:
		localctx = NewFunctionCallExpressionAtomContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(464)
			p.FunctionCall()
		}


	case 4:
		localctx = NewUnaryExpressionAtomContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(465)
			p.UnaryOperator()
		}
		{
			p.SetState(466)
			p.expressionAtom(6)
		}


	case 5:
		localctx = NewNestedExpressionAtomContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(468)
			p.Match(MySqlParserLR_BRACKET)
		}
		{
			p.SetState(469)
			p.expression(0)
		}
		p.SetState(474)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for _la == MySqlParserCOMMA {
			{
				p.SetState(470)
				p.Match(MySqlParserCOMMA)
			}
			{
				p.SetState(471)
				p.expression(0)
			}


			p.SetState(476)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(477)
			p.Match(MySqlParserRR_BRACKET)
		}


	case 6:
		localctx = NewExistsExpressionAtomContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(479)
			p.Match(MySqlParserEXISTS)
		}
		{
			p.SetState(480)
			p.Match(MySqlParserLR_BRACKET)
		}
		{
			p.SetState(481)
			p.SelectStatement()
		}
		{
			p.SetState(482)
			p.Match(MySqlParserRR_BRACKET)
		}


	case 7:
		localctx = NewSubqueryExpressionAtomContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(484)
			p.Match(MySqlParserLR_BRACKET)
		}
		{
			p.SetState(485)
			p.SelectStatement()
		}
		{
			p.SetState(486)
			p.Match(MySqlParserRR_BRACKET)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(499)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 68, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(497)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 67, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPriorityMathExpressionAtomContext(p, NewExpressionAtomContext(p, _parentctx, _parentState))
				localctx.(*PriorityMathExpressionAtomContext).left = _prevctx


				p.PushNewRecursionContext(localctx, _startState, MySqlParserRULE_expressionAtom)
				p.SetState(490)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(491)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*PriorityMathExpressionAtomContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((((_la - 58)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 58))) & ((1 << (MySqlParserSTAR - 58)) | (1 << (MySqlParserDIVIDE - 58)) | (1 << (MySqlParserMODULE - 58)) | (1 << (MySqlParserDIV - 58)) | (1 << (MySqlParserMOD - 58)))) != 0)) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*PriorityMathExpressionAtomContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(492)

					var _x = p.expressionAtom(3)

					localctx.(*PriorityMathExpressionAtomContext).right = _x
				}


			case 2:
				localctx = NewMathExpressionAtomContext(p, NewExpressionAtomContext(p, _parentctx, _parentState))
				localctx.(*MathExpressionAtomContext).left = _prevctx


				p.PushNewRecursionContext(localctx, _startState, MySqlParserRULE_expressionAtom)
				p.SetState(493)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(494)
					p.MathOperator()
				}
				{
					p.SetState(495)

					var _x = p.expressionAtom(2)

					localctx.(*MathExpressionAtomContext).right = _x
				}

			}

		}
		p.SetState(501)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 68, p.GetParserRuleContext())
	}



	return localctx
}


// IUnaryOperatorContext is an interface to support dynamic dispatch.
type IUnaryOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnaryOperatorContext differentiates from other interfaces.
	IsUnaryOperatorContext()
}

type UnaryOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryOperatorContext() *UnaryOperatorContext {
	var p = new(UnaryOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySqlParserRULE_unaryOperator
	return p
}

func (*UnaryOperatorContext) IsUnaryOperatorContext() {}

func NewUnaryOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryOperatorContext {
	var p = new(UnaryOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySqlParserRULE_unaryOperator

	return p
}

func (s *UnaryOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryOperatorContext) EXCLAMATION_SYMBOL() antlr.TerminalNode {
	return s.GetToken(MySqlParserEXCLAMATION_SYMBOL, 0)
}

func (s *UnaryOperatorContext) BIT_NOT_OP() antlr.TerminalNode {
	return s.GetToken(MySqlParserBIT_NOT_OP, 0)
}

func (s *UnaryOperatorContext) PLUS() antlr.TerminalNode {
	return s.GetToken(MySqlParserPLUS, 0)
}

func (s *UnaryOperatorContext) MINUS() antlr.TerminalNode {
	return s.GetToken(MySqlParserMINUS, 0)
}

func (s *UnaryOperatorContext) NOT() antlr.TerminalNode {
	return s.GetToken(MySqlParserNOT, 0)
}

func (s *UnaryOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *UnaryOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterUnaryOperator(s)
	}
}

func (s *UnaryOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitUnaryOperator(s)
	}
}

func (s *UnaryOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitUnaryOperator(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MySqlParser) UnaryOperator() (localctx IUnaryOperatorContext) {
	localctx = NewUnaryOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, MySqlParserRULE_unaryOperator)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(502)
		_la = p.GetTokenStream().LA(1)

		if !(_la == MySqlParserNOT || ((((_la - 61)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 61))) & ((1 << (MySqlParserPLUS - 61)) | (1 << (MySqlParserMINUS - 61)) | (1 << (MySqlParserEXCLAMATION_SYMBOL - 61)) | (1 << (MySqlParserBIT_NOT_OP - 61)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


// IComparisonOperatorContext is an interface to support dynamic dispatch.
type IComparisonOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComparisonOperatorContext differentiates from other interfaces.
	IsComparisonOperatorContext()
}

type ComparisonOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonOperatorContext() *ComparisonOperatorContext {
	var p = new(ComparisonOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySqlParserRULE_comparisonOperator
	return p
}

func (*ComparisonOperatorContext) IsComparisonOperatorContext() {}

func NewComparisonOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonOperatorContext {
	var p = new(ComparisonOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySqlParserRULE_comparisonOperator

	return p
}

func (s *ComparisonOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonOperatorContext) EQUAL_SYMBOL() antlr.TerminalNode {
	return s.GetToken(MySqlParserEQUAL_SYMBOL, 0)
}

func (s *ComparisonOperatorContext) GREATER_SYMBOL() antlr.TerminalNode {
	return s.GetToken(MySqlParserGREATER_SYMBOL, 0)
}

func (s *ComparisonOperatorContext) LESS_SYMBOL() antlr.TerminalNode {
	return s.GetToken(MySqlParserLESS_SYMBOL, 0)
}

func (s *ComparisonOperatorContext) EXCLAMATION_SYMBOL() antlr.TerminalNode {
	return s.GetToken(MySqlParserEXCLAMATION_SYMBOL, 0)
}

func (s *ComparisonOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ComparisonOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterComparisonOperator(s)
	}
}

func (s *ComparisonOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitComparisonOperator(s)
	}
}

func (s *ComparisonOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitComparisonOperator(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MySqlParser) ComparisonOperator() (localctx IComparisonOperatorContext) {
	localctx = NewComparisonOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, MySqlParserRULE_comparisonOperator)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(518)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 69, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(504)
			p.Match(MySqlParserEQUAL_SYMBOL)
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(505)
			p.Match(MySqlParserGREATER_SYMBOL)
		}


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(506)
			p.Match(MySqlParserLESS_SYMBOL)
		}


	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(507)
			p.Match(MySqlParserLESS_SYMBOL)
		}
		{
			p.SetState(508)
			p.Match(MySqlParserEQUAL_SYMBOL)
		}


	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(509)
			p.Match(MySqlParserGREATER_SYMBOL)
		}
		{
			p.SetState(510)
			p.Match(MySqlParserEQUAL_SYMBOL)
		}


	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(511)
			p.Match(MySqlParserLESS_SYMBOL)
		}
		{
			p.SetState(512)
			p.Match(MySqlParserGREATER_SYMBOL)
		}


	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(513)
			p.Match(MySqlParserEXCLAMATION_SYMBOL)
		}
		{
			p.SetState(514)
			p.Match(MySqlParserEQUAL_SYMBOL)
		}


	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(515)
			p.Match(MySqlParserLESS_SYMBOL)
		}
		{
			p.SetState(516)
			p.Match(MySqlParserEQUAL_SYMBOL)
		}
		{
			p.SetState(517)
			p.Match(MySqlParserGREATER_SYMBOL)
		}

	}


	return localctx
}


// ILogicalOperatorContext is an interface to support dynamic dispatch.
type ILogicalOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLogicalOperatorContext differentiates from other interfaces.
	IsLogicalOperatorContext()
}

type LogicalOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicalOperatorContext() *LogicalOperatorContext {
	var p = new(LogicalOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySqlParserRULE_logicalOperator
	return p
}

func (*LogicalOperatorContext) IsLogicalOperatorContext() {}

func NewLogicalOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalOperatorContext {
	var p = new(LogicalOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySqlParserRULE_logicalOperator

	return p
}

func (s *LogicalOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalOperatorContext) AND() antlr.TerminalNode {
	return s.GetToken(MySqlParserAND, 0)
}

func (s *LogicalOperatorContext) OR() antlr.TerminalNode {
	return s.GetToken(MySqlParserOR, 0)
}

func (s *LogicalOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *LogicalOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterLogicalOperator(s)
	}
}

func (s *LogicalOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitLogicalOperator(s)
	}
}

func (s *LogicalOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitLogicalOperator(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MySqlParser) LogicalOperator() (localctx ILogicalOperatorContext) {
	localctx = NewLogicalOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, MySqlParserRULE_logicalOperator)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(520)
		_la = p.GetTokenStream().LA(1)

		if !(_la == MySqlParserAND || _la == MySqlParserOR) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


// IMathOperatorContext is an interface to support dynamic dispatch.
type IMathOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMathOperatorContext differentiates from other interfaces.
	IsMathOperatorContext()
}

type MathOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMathOperatorContext() *MathOperatorContext {
	var p = new(MathOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySqlParserRULE_mathOperator
	return p
}

func (*MathOperatorContext) IsMathOperatorContext() {}

func NewMathOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MathOperatorContext {
	var p = new(MathOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySqlParserRULE_mathOperator

	return p
}

func (s *MathOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *MathOperatorContext) PLUS() antlr.TerminalNode {
	return s.GetToken(MySqlParserPLUS, 0)
}

func (s *MathOperatorContext) MINUS() antlr.TerminalNode {
	return s.GetToken(MySqlParserMINUS, 0)
}

func (s *MathOperatorContext) MINUSMINUS() antlr.TerminalNode {
	return s.GetToken(MySqlParserMINUSMINUS, 0)
}

func (s *MathOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MathOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *MathOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterMathOperator(s)
	}
}

func (s *MathOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitMathOperator(s)
	}
}

func (s *MathOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitMathOperator(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MySqlParser) MathOperator() (localctx IMathOperatorContext) {
	localctx = NewMathOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, MySqlParserRULE_mathOperator)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(522)
		_la = p.GetTokenStream().LA(1)

		if !(((((_la - 61)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 61))) & ((1 << (MySqlParserPLUS - 61)) | (1 << (MySqlParserMINUSMINUS - 61)) | (1 << (MySqlParserMINUS - 61)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


// IInsertStatementContext is an interface to support dynamic dispatch.
type IInsertStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetColumns returns the columns rule contexts.
	GetColumns() IUidListContext


	// SetColumns sets the columns rule contexts.
	SetColumns(IUidListContext)


	// IsInsertStatementContext differentiates from other interfaces.
	IsInsertStatementContext()
}

type InsertStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	columns IUidListContext 
}

func NewEmptyInsertStatementContext() *InsertStatementContext {
	var p = new(InsertStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySqlParserRULE_insertStatement
	return p
}

func (*InsertStatementContext) IsInsertStatementContext() {}

func NewInsertStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InsertStatementContext {
	var p = new(InsertStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySqlParserRULE_insertStatement

	return p
}

func (s *InsertStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *InsertStatementContext) GetColumns() IUidListContext { return s.columns }


func (s *InsertStatementContext) SetColumns(v IUidListContext) { s.columns = v }


func (s *InsertStatementContext) INSERT() antlr.TerminalNode {
	return s.GetToken(MySqlParserINSERT, 0)
}

func (s *InsertStatementContext) ID() antlr.TerminalNode {
	return s.GetToken(MySqlParserID, 0)
}

func (s *InsertStatementContext) InsertStatementValue() IInsertStatementValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInsertStatementValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInsertStatementValueContext)
}

func (s *InsertStatementContext) IGNORE() antlr.TerminalNode {
	return s.GetToken(MySqlParserIGNORE, 0)
}

func (s *InsertStatementContext) INTO() antlr.TerminalNode {
	return s.GetToken(MySqlParserINTO, 0)
}

func (s *InsertStatementContext) LR_BRACKET() antlr.TerminalNode {
	return s.GetToken(MySqlParserLR_BRACKET, 0)
}

func (s *InsertStatementContext) RR_BRACKET() antlr.TerminalNode {
	return s.GetToken(MySqlParserRR_BRACKET, 0)
}

func (s *InsertStatementContext) UidList() IUidListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUidListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUidListContext)
}

func (s *InsertStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InsertStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *InsertStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterInsertStatement(s)
	}
}

func (s *InsertStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitInsertStatement(s)
	}
}

func (s *InsertStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitInsertStatement(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MySqlParser) InsertStatement() (localctx IInsertStatementContext) {
	localctx = NewInsertStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, MySqlParserRULE_insertStatement)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(524)
		p.Match(MySqlParserINSERT)
	}
	p.SetState(526)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == MySqlParserIGNORE {
		{
			p.SetState(525)
			p.Match(MySqlParserIGNORE)
		}

	}
	p.SetState(529)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == MySqlParserINTO {
		{
			p.SetState(528)
			p.Match(MySqlParserINTO)
		}

	}
	{
		p.SetState(531)
		p.Match(MySqlParserID)
	}
	p.SetState(536)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 72, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(532)
			p.Match(MySqlParserLR_BRACKET)
		}
		{
			p.SetState(533)

			var _x = p.UidList()


			localctx.(*InsertStatementContext).columns = _x
		}
		{
			p.SetState(534)
			p.Match(MySqlParserRR_BRACKET)
		}


	}
	{
		p.SetState(538)
		p.InsertStatementValue()
	}



	return localctx
}


// IInsertStatementValueContext is an interface to support dynamic dispatch.
type IInsertStatementValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetInsertFormat returns the insertFormat token.
	GetInsertFormat() antlr.Token 


	// SetInsertFormat sets the insertFormat token.
	SetInsertFormat(antlr.Token) 


	// IsInsertStatementValueContext differentiates from other interfaces.
	IsInsertStatementValueContext()
}

type InsertStatementValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	insertFormat antlr.Token
}

func NewEmptyInsertStatementValueContext() *InsertStatementValueContext {
	var p = new(InsertStatementValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySqlParserRULE_insertStatementValue
	return p
}

func (*InsertStatementValueContext) IsInsertStatementValueContext() {}

func NewInsertStatementValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InsertStatementValueContext {
	var p = new(InsertStatementValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySqlParserRULE_insertStatementValue

	return p
}

func (s *InsertStatementValueContext) GetParser() antlr.Parser { return s.parser }

func (s *InsertStatementValueContext) GetInsertFormat() antlr.Token { return s.insertFormat }


func (s *InsertStatementValueContext) SetInsertFormat(v antlr.Token) { s.insertFormat = v }


func (s *InsertStatementValueContext) SelectStatement() ISelectStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectStatementContext)
}

func (s *InsertStatementValueContext) AllLR_BRACKET() []antlr.TerminalNode {
	return s.GetTokens(MySqlParserLR_BRACKET)
}

func (s *InsertStatementValueContext) LR_BRACKET(i int) antlr.TerminalNode {
	return s.GetToken(MySqlParserLR_BRACKET, i)
}

func (s *InsertStatementValueContext) AllExpressionsWithDefaults() []IExpressionsWithDefaultsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionsWithDefaultsContext)(nil)).Elem())
	var tst = make([]IExpressionsWithDefaultsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionsWithDefaultsContext)
		}
	}

	return tst
}

func (s *InsertStatementValueContext) ExpressionsWithDefaults(i int) IExpressionsWithDefaultsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionsWithDefaultsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionsWithDefaultsContext)
}

func (s *InsertStatementValueContext) AllRR_BRACKET() []antlr.TerminalNode {
	return s.GetTokens(MySqlParserRR_BRACKET)
}

func (s *InsertStatementValueContext) RR_BRACKET(i int) antlr.TerminalNode {
	return s.GetToken(MySqlParserRR_BRACKET, i)
}

func (s *InsertStatementValueContext) VALUES() antlr.TerminalNode {
	return s.GetToken(MySqlParserVALUES, 0)
}

func (s *InsertStatementValueContext) VALUE() antlr.TerminalNode {
	return s.GetToken(MySqlParserVALUE, 0)
}

func (s *InsertStatementValueContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MySqlParserCOMMA)
}

func (s *InsertStatementValueContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MySqlParserCOMMA, i)
}

func (s *InsertStatementValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InsertStatementValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *InsertStatementValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterInsertStatementValue(s)
	}
}

func (s *InsertStatementValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitInsertStatementValue(s)
	}
}

func (s *InsertStatementValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitInsertStatementValue(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MySqlParser) InsertStatementValue() (localctx IInsertStatementValueContext) {
	localctx = NewInsertStatementValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, MySqlParserRULE_insertStatementValue)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(555)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MySqlParserSELECT, MySqlParserLR_BRACKET:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(540)
			p.SelectStatement()
		}


	case MySqlParserVALUES, MySqlParserVALUE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(541)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*InsertStatementValueContext).insertFormat = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == MySqlParserVALUES || _la == MySqlParserVALUE) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*InsertStatementValueContext).insertFormat = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(542)
			p.Match(MySqlParserLR_BRACKET)
		}
		{
			p.SetState(543)
			p.ExpressionsWithDefaults()
		}
		{
			p.SetState(544)
			p.Match(MySqlParserRR_BRACKET)
		}
		p.SetState(552)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for _la == MySqlParserCOMMA {
			{
				p.SetState(545)
				p.Match(MySqlParserCOMMA)
			}
			{
				p.SetState(546)
				p.Match(MySqlParserLR_BRACKET)
			}
			{
				p.SetState(547)
				p.ExpressionsWithDefaults()
			}
			{
				p.SetState(548)
				p.Match(MySqlParserRR_BRACKET)
			}


			p.SetState(554)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IUpdatedElementContext is an interface to support dynamic dispatch.
type IUpdatedElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUpdatedElementContext differentiates from other interfaces.
	IsUpdatedElementContext()
}

type UpdatedElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpdatedElementContext() *UpdatedElementContext {
	var p = new(UpdatedElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySqlParserRULE_updatedElement
	return p
}

func (*UpdatedElementContext) IsUpdatedElementContext() {}

func NewUpdatedElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UpdatedElementContext {
	var p = new(UpdatedElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySqlParserRULE_updatedElement

	return p
}

func (s *UpdatedElementContext) GetParser() antlr.Parser { return s.parser }

func (s *UpdatedElementContext) FullColumnName() IFullColumnNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFullColumnNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFullColumnNameContext)
}

func (s *UpdatedElementContext) EQUAL_SYMBOL() antlr.TerminalNode {
	return s.GetToken(MySqlParserEQUAL_SYMBOL, 0)
}

func (s *UpdatedElementContext) ExpressionOrDefault() IExpressionOrDefaultContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionOrDefaultContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionOrDefaultContext)
}

func (s *UpdatedElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UpdatedElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *UpdatedElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterUpdatedElement(s)
	}
}

func (s *UpdatedElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitUpdatedElement(s)
	}
}

func (s *UpdatedElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitUpdatedElement(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MySqlParser) UpdatedElement() (localctx IUpdatedElementContext) {
	localctx = NewUpdatedElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, MySqlParserRULE_updatedElement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(557)
		p.FullColumnName()
	}
	{
		p.SetState(558)
		p.Match(MySqlParserEQUAL_SYMBOL)
	}
	{
		p.SetState(559)
		p.ExpressionOrDefault()
	}



	return localctx
}


// IExpressionsWithDefaultsContext is an interface to support dynamic dispatch.
type IExpressionsWithDefaultsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionsWithDefaultsContext differentiates from other interfaces.
	IsExpressionsWithDefaultsContext()
}

type ExpressionsWithDefaultsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionsWithDefaultsContext() *ExpressionsWithDefaultsContext {
	var p = new(ExpressionsWithDefaultsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySqlParserRULE_expressionsWithDefaults
	return p
}

func (*ExpressionsWithDefaultsContext) IsExpressionsWithDefaultsContext() {}

func NewExpressionsWithDefaultsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionsWithDefaultsContext {
	var p = new(ExpressionsWithDefaultsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySqlParserRULE_expressionsWithDefaults

	return p
}

func (s *ExpressionsWithDefaultsContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionsWithDefaultsContext) AllExpressionOrDefault() []IExpressionOrDefaultContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionOrDefaultContext)(nil)).Elem())
	var tst = make([]IExpressionOrDefaultContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionOrDefaultContext)
		}
	}

	return tst
}

func (s *ExpressionsWithDefaultsContext) ExpressionOrDefault(i int) IExpressionOrDefaultContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionOrDefaultContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionOrDefaultContext)
}

func (s *ExpressionsWithDefaultsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MySqlParserCOMMA)
}

func (s *ExpressionsWithDefaultsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MySqlParserCOMMA, i)
}

func (s *ExpressionsWithDefaultsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionsWithDefaultsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ExpressionsWithDefaultsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterExpressionsWithDefaults(s)
	}
}

func (s *ExpressionsWithDefaultsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitExpressionsWithDefaults(s)
	}
}

func (s *ExpressionsWithDefaultsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitExpressionsWithDefaults(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MySqlParser) ExpressionsWithDefaults() (localctx IExpressionsWithDefaultsContext) {
	localctx = NewExpressionsWithDefaultsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, MySqlParserRULE_expressionsWithDefaults)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(561)
		p.ExpressionOrDefault()
	}
	p.SetState(566)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == MySqlParserCOMMA {
		{
			p.SetState(562)
			p.Match(MySqlParserCOMMA)
		}
		{
			p.SetState(563)
			p.ExpressionOrDefault()
		}


		p.SetState(568)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


// IExpressionOrDefaultContext is an interface to support dynamic dispatch.
type IExpressionOrDefaultContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionOrDefaultContext differentiates from other interfaces.
	IsExpressionOrDefaultContext()
}

type ExpressionOrDefaultContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionOrDefaultContext() *ExpressionOrDefaultContext {
	var p = new(ExpressionOrDefaultContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySqlParserRULE_expressionOrDefault
	return p
}

func (*ExpressionOrDefaultContext) IsExpressionOrDefaultContext() {}

func NewExpressionOrDefaultContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionOrDefaultContext {
	var p = new(ExpressionOrDefaultContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySqlParserRULE_expressionOrDefault

	return p
}

func (s *ExpressionOrDefaultContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionOrDefaultContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionOrDefaultContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(MySqlParserDEFAULT, 0)
}

func (s *ExpressionOrDefaultContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionOrDefaultContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ExpressionOrDefaultContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterExpressionOrDefault(s)
	}
}

func (s *ExpressionOrDefaultContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitExpressionOrDefault(s)
	}
}

func (s *ExpressionOrDefaultContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitExpressionOrDefault(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MySqlParser) ExpressionOrDefault() (localctx IExpressionOrDefaultContext) {
	localctx = NewExpressionOrDefaultContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, MySqlParserRULE_expressionOrDefault)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(571)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MySqlParserCASE, MySqlParserEXISTS, MySqlParserFALSE, MySqlParserNOT, MySqlParserNULL_LITERAL, MySqlParserTRUE, MySqlParserPLUS, MySqlParserMINUS, MySqlParserEXCLAMATION_SYMBOL, MySqlParserBIT_NOT_OP, MySqlParserLR_BRACKET, MySqlParserSTRING_LITERAL, MySqlParserDECIMAL_LITERAL, MySqlParserREAL_LITERAL, MySqlParserNULL_SPEC_LITERAL, MySqlParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(569)
			p.expression(0)
		}


	case MySqlParserDEFAULT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(570)
			p.Match(MySqlParserDEFAULT)
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IDeleteStatementContext is an interface to support dynamic dispatch.
type IDeleteStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeleteStatementContext differentiates from other interfaces.
	IsDeleteStatementContext()
}

type DeleteStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeleteStatementContext() *DeleteStatementContext {
	var p = new(DeleteStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySqlParserRULE_deleteStatement
	return p
}

func (*DeleteStatementContext) IsDeleteStatementContext() {}

func NewDeleteStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeleteStatementContext {
	var p = new(DeleteStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySqlParserRULE_deleteStatement

	return p
}

func (s *DeleteStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *DeleteStatementContext) DELETE() antlr.TerminalNode {
	return s.GetToken(MySqlParserDELETE, 0)
}

func (s *DeleteStatementContext) FROM() antlr.TerminalNode {
	return s.GetToken(MySqlParserFROM, 0)
}

func (s *DeleteStatementContext) ID() antlr.TerminalNode {
	return s.GetToken(MySqlParserID, 0)
}

func (s *DeleteStatementContext) IGNORE() antlr.TerminalNode {
	return s.GetToken(MySqlParserIGNORE, 0)
}

func (s *DeleteStatementContext) WHERE() antlr.TerminalNode {
	return s.GetToken(MySqlParserWHERE, 0)
}

func (s *DeleteStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DeleteStatementContext) OrderByClause() IOrderByClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrderByClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOrderByClauseContext)
}

func (s *DeleteStatementContext) LimitClause() ILimitClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILimitClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILimitClauseContext)
}

func (s *DeleteStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeleteStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *DeleteStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterDeleteStatement(s)
	}
}

func (s *DeleteStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitDeleteStatement(s)
	}
}

func (s *DeleteStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitDeleteStatement(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MySqlParser) DeleteStatement() (localctx IDeleteStatementContext) {
	localctx = NewDeleteStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, MySqlParserRULE_deleteStatement)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(573)
		p.Match(MySqlParserDELETE)
	}
	p.SetState(575)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == MySqlParserIGNORE {
		{
			p.SetState(574)
			p.Match(MySqlParserIGNORE)
		}

	}
	{
		p.SetState(577)
		p.Match(MySqlParserFROM)
	}
	{
		p.SetState(578)
		p.Match(MySqlParserID)
	}
	p.SetState(581)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == MySqlParserWHERE {
		{
			p.SetState(579)
			p.Match(MySqlParserWHERE)
		}
		{
			p.SetState(580)
			p.expression(0)
		}

	}
	p.SetState(584)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == MySqlParserORDER {
		{
			p.SetState(583)
			p.OrderByClause()
		}

	}
	p.SetState(587)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == MySqlParserLIMIT {
		{
			p.SetState(586)
			p.LimitClause()
		}

	}



	return localctx
}


// IUpdateStatementContext is an interface to support dynamic dispatch.
type IUpdateStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUpdateStatementContext differentiates from other interfaces.
	IsUpdateStatementContext()
}

type UpdateStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpdateStatementContext() *UpdateStatementContext {
	var p = new(UpdateStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MySqlParserRULE_updateStatement
	return p
}

func (*UpdateStatementContext) IsUpdateStatementContext() {}

func NewUpdateStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UpdateStatementContext {
	var p = new(UpdateStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySqlParserRULE_updateStatement

	return p
}

func (s *UpdateStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *UpdateStatementContext) UPDATE() antlr.TerminalNode {
	return s.GetToken(MySqlParserUPDATE, 0)
}

func (s *UpdateStatementContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(MySqlParserID)
}

func (s *UpdateStatementContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(MySqlParserID, i)
}

func (s *UpdateStatementContext) SET() antlr.TerminalNode {
	return s.GetToken(MySqlParserSET, 0)
}

func (s *UpdateStatementContext) AllUpdatedElement() []IUpdatedElementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IUpdatedElementContext)(nil)).Elem())
	var tst = make([]IUpdatedElementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IUpdatedElementContext)
		}
	}

	return tst
}

func (s *UpdateStatementContext) UpdatedElement(i int) IUpdatedElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUpdatedElementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IUpdatedElementContext)
}

func (s *UpdateStatementContext) IGNORE() antlr.TerminalNode {
	return s.GetToken(MySqlParserIGNORE, 0)
}

func (s *UpdateStatementContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MySqlParserCOMMA)
}

func (s *UpdateStatementContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MySqlParserCOMMA, i)
}

func (s *UpdateStatementContext) WHERE() antlr.TerminalNode {
	return s.GetToken(MySqlParserWHERE, 0)
}

func (s *UpdateStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *UpdateStatementContext) OrderByClause() IOrderByClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrderByClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOrderByClauseContext)
}

func (s *UpdateStatementContext) LimitClause() ILimitClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILimitClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILimitClauseContext)
}

func (s *UpdateStatementContext) AS() antlr.TerminalNode {
	return s.GetToken(MySqlParserAS, 0)
}

func (s *UpdateStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UpdateStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *UpdateStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.EnterUpdateStatement(s)
	}
}

func (s *UpdateStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySqlParserListener); ok {
		listenerT.ExitUpdateStatement(s)
	}
}

func (s *UpdateStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case MySqlParserVisitor:
		return t.VisitUpdateStatement(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *MySqlParser) UpdateStatement() (localctx IUpdateStatementContext) {
	localctx = NewUpdateStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, MySqlParserRULE_updateStatement)
	var _la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(589)
		p.Match(MySqlParserUPDATE)
	}
	p.SetState(591)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == MySqlParserIGNORE {
		{
			p.SetState(590)
			p.Match(MySqlParserIGNORE)
		}

	}
	{
		p.SetState(593)
		p.Match(MySqlParserID)
	}
	p.SetState(598)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == MySqlParserAS || _la == MySqlParserID {
		p.SetState(595)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		if _la == MySqlParserAS {
			{
				p.SetState(594)
				p.Match(MySqlParserAS)
			}

		}
		{
			p.SetState(597)
			p.Match(MySqlParserID)
		}

	}
	{
		p.SetState(600)
		p.Match(MySqlParserSET)
	}
	{
		p.SetState(601)
		p.UpdatedElement()
	}
	p.SetState(606)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == MySqlParserCOMMA {
		{
			p.SetState(602)
			p.Match(MySqlParserCOMMA)
		}
		{
			p.SetState(603)
			p.UpdatedElement()
		}


		p.SetState(608)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(611)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == MySqlParserWHERE {
		{
			p.SetState(609)
			p.Match(MySqlParserWHERE)
		}
		{
			p.SetState(610)
			p.expression(0)
		}

	}
	p.SetState(614)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == MySqlParserORDER {
		{
			p.SetState(613)
			p.OrderByClause()
		}

	}
	p.SetState(617)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == MySqlParserLIMIT {
		{
			p.SetState(616)
			p.LimitClause()
		}

	}



	return localctx
}


func (p *MySqlParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 27:
			var t *ExpressionContext = nil
			if localctx != nil { t = localctx.(*ExpressionContext) }
			return p.Expression_Sempred(t, predIndex)

	case 28:
			var t *PredicateContext = nil
			if localctx != nil { t = localctx.(*PredicateContext) }
			return p.Predicate_Sempred(t, predIndex)

	case 29:
			var t *ExpressionAtomContext = nil
			if localctx != nil { t = localctx.(*ExpressionAtomContext) }
			return p.ExpressionAtom_Sempred(t, predIndex)


	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *MySqlParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
			return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *MySqlParser) Predicate_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
			return p.Precpred(p.GetParserRuleContext(), 6)

	case 2:
			return p.Precpred(p.GetParserRuleContext(), 4)

	case 3:
			return p.Precpred(p.GetParserRuleContext(), 3)

	case 4:
			return p.Precpred(p.GetParserRuleContext(), 2)

	case 5:
			return p.Precpred(p.GetParserRuleContext(), 8)

	case 6:
			return p.Precpred(p.GetParserRuleContext(), 7)

	case 7:
			return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *MySqlParser) ExpressionAtom_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 8:
			return p.Precpred(p.GetParserRuleContext(), 2)

	case 9:
			return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

