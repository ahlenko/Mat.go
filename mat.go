// Бібліотека для виконання арифметичних операцій над матрицями
//
// Розробники:
//  Ахленко Д.А.
//  Талибов Е.Т.
// Дата редагування:
//  02.04.2022
package Mat

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Структура Matrix:
//  uint g_size - стовпці
//  uint v_size - рядки
//  []float32 elements - елементи матриці
type Matrix struct {
	g_size   uint
	v_size   uint
	elements []float32
}

// (Privat) Функція створення матриці
//
// Вхідні дані:
//  uint x - кількість рядків
//  uint y - кількість стовпців
// Вихідні дані:
//  make Matrix - структура Matrix
func setMatrix(x, y uint) Matrix {
	var mat Matrix
	mat.v_size = x
	mat.g_size = y
	mat.elements = make([]float32, x*y)
	return mat
}

// (Privat) Метод перетворення масиву до матриці
//
// Вхідні дані:
//  [][]float32 - матриця у форматі двомірного масиву
// Вихідні дані:
//  make Matrix - структура Matrix
func (mat *Matrix) toArray() [][]float32 {
	res := make([][]float32, mat.v_size)
	for i := range res {
		res[i] = make([]float32, mat.g_size)
	}
	el := 0
	for i := uint(0); i < mat.g_size; i++ {
		for j := uint(0); j < mat.v_size; j++ {
			res[i][j] = mat.elements[el]
			el++
		}
	}
	return res
}

// (Privat) Метод перетворення матриці до масиву
//
// Вхідні дані:
//  make Matrix - структура Matrix
// Вихідні дані:
//  [][]float32 - матриця у форматі двомірного масиву
func (mat *Matrix) toMatrix(inpt [][]float32) Matrix {
	res := setMatrix(mat.g_size, mat.v_size)
	el := uint(0)
	for i := uint(0); i < mat.g_size; i++ {
		for j := uint(0); j < mat.v_size; j++ {
			res.elements[el] = inpt[i][j]
			el++
		}
	}
	return res
}

// Метод копіювання матриці
//
// Вхідні дані:
//  * Matrix - структура Matrix
// Вихідні дані:
//  make Matrix - структура Matrix
func (mat *Matrix) MatrixCopy() Matrix {
	outp := setMatrix(mat.g_size, mat.v_size)
	for i := range mat.elements {
		outp.elements[i] = mat.elements[i]
	}
	return outp
}

// Функція заповнення матриці з одномірного масиву
//  - Якщо кількість елементів у масиві менша за введену розмірність, надлишкові елементи заповнюються нулями;
//  - Якщо введена розмірність = 0 то повертається порожня матриця;
// Вхідні дані:
//  []float32 inpt - одномірний масив
//  uint size_g - кількість стовпців
//  uint size_v - кількість рядків
// Вихідні дані:
//  make Matrix - структура Matrix
func ProgReader(inpt []float32, size_g, size_v uint) Matrix {
	mat := setMatrix(size_g, size_v)
	if size_g == 0 || size_v == 0 {
		return mat
	}
	size := 0
	for i := range inpt {
		size = i
	}
	if size < int(size_g*size_v) {
		el := 0
		for i := range inpt {
			mat.elements[i] = inpt[i]
			el++
		}
		for el < int(size_g*size_v) {
			mat.elements[el] = 0
			el++
		}
	} else {
		for i := 0; i < int(size_g*size_v); i++ {
			mat.elements[i] = inpt[i]
		}
	}
	return mat
}

// Функція заповнення матриці з клавіатури
//  - Вводиться розмірність матриці(значення перевіряються);
//  - Вводяться елементи матриці;
// Вихідні дані:
//  make Matrix - структура Matrix
func KeyFilling() Matrix {
	var size_g, size_v int
	fmt.Println("Введіть розмірність матриці:")
	for size_g == 0 {
		fmt.Print(":стовпці -> ")
		fmt.Scanln(&size_g)
		if size_g <= 0 {
			fmt.Println("Помилка введення")
		}
	}
	for size_v == 0 {
		fmt.Print(":рядки -> ")
		fmt.Scanln(&size_v)
		if size_v <= 0 {
			fmt.Println("Помилка введення")
		}
	}
	mat := setMatrix(uint(size_g), uint(size_v))
	y := make([][]float32, size_v)
	for i := range y {
		y[i] = make([]float32, size_g)
	}
	fmt.Println("Введіть елементи матриці:")
	for i := 0; i < size_v; i++ {
		for j := 0; j < size_g; j++ {
			fmt.Print("El[", i, "][", j, "] = ")
			fmt.Scanln(&y[i][j])
		}
	}
	return mat.toMatrix(y)
}

// Функція заповнення матриці з клавіатури
//  - Вводяться елементи матриці;
// Вхідні дані:
//  uint size_g - кількість стовпців
//  uint size_v - кількість рядків
// Вихідні дані:
//  make Matrix - структура Matrix
func KeySizeFilling(size_g, size_v uint) Matrix {
	mat := setMatrix(size_g, size_v)
	y := make([][]float32, size_v)
	for i := range y {
		y[i] = make([]float32, size_g)
	}
	if size_g == 0 || size_v == 0 {
		return mat
	}
	fmt.Println("Введіть елементи матриці:")
	for i := uint(0); i < size_v; i++ {
		for j := uint(0); j < size_g; j++ {
			fmt.Print("El[", i, "][", j, "] = ")
			fmt.Scanln(&y[i][j])
		}
	}
	return mat.toMatrix(y)
}

// Функція читання матриці з файлу
//
// Вхідні дані:
//  string filename - назва файлу
// Вихідні дані:
//  make Matrix - структура Matrix
//  int:
//   -1 - помилка роботи з файлом
//   0 - помилок немає
func FileReader(filename string) (Matrix, int) {
	mat := setMatrix(0, 0)
	var temp []float32
	var num float64
	var num_32 float32
	var size_g, size_v uint

	file, err := os.Open(filename)
	if err != nil {
		return mat, -1
	}
	i := bufio.NewReader(file)
	el := 0

	for {
		buff, _, err := i.ReadLine() // Читаем по строке
		if err != nil {
			if err == io.EOF { // Файл закончился
				break
			} else {
				return mat, -1
			}
		}
		size_g = 0
		buff = append(buff, 32)
		var add float32
		i := 0
		for i < len(buff) {
			add = 1
			for buff[i] != 46 {
				num, _ = strconv.ParseFloat(string(buff[i]), 32)
				num_32 = float32(num)
				temp = append(temp, 0)
				temp[el] *= add
				temp[el] += num_32
				add *= 10
				i++
			}
			i++
			add = 10
			for buff[i] != 32 {
				num, _ = strconv.ParseFloat(string(buff[i]), 32)
				num_32 = float32(num)
				temp[el] += num_32 / add
				add *= 10
				i++

			}
			size_g++
			i++
			el++
		}
		size_v++
	}
	file.Close()
	mat = ProgReader(temp, size_g, size_v)
	return mat, 0
}

// Функція заповнення матриці випадковими значеннями
//  - Якщо введена розмірність = 0, повертається порожня матриця;
//  - Якщо значення max,min = 0, матриця одинична;
// Вхідні дані:
//  uint size_g - кількість стовпців
//  uint size_v - кількість рядків
//  float max - максимальне значення
//  float min - мінімальне значення
// Вихідні дані:
//  make Matrix - структура Matrix
func RandMatrix(size_g, size_v uint, min, max float32) Matrix {
	mat := setMatrix(size_g, size_v)
	if size_g == 0 || size_v == 0 {
		return mat
	} else if min == 0 && max == 0 {
		for i := uint(0); i < size_g*size_v; i += size_g + 1 {
			mat.elements[i] = 1
		}
	} else {
		for i := range mat.elements {
			mat.elements[i] = min + rand.Float32()*(max-min)
		}
	}
	return mat
}

// Метод заміни елемента матриці
//  - змінює елемент за порядковим номером;
// Вхідні дані:
//  uint adr - номер елементу
//  float rep - нове значення
// Вихідні дані:
//  *Matrix - структура Matrix
func (mat *Matrix) ReplausElem(adr uint, rep float32) int {
	if adr > (mat.g_size * mat.v_size) {
		return -1
	}
	mat.elements[adr] = rep
	return 0
}

// Метод запису матриці у файл
//  - імя - матриця - час запису;
// Вхідні дані:
//  string matname - назва матриці
//  string filename - назва файлу
func (mat *Matrix) FileWriter(matname, filename string) {
	var st string
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		file, _ = os.Create(filename)
	}
	file.WriteString("\n" + matname + "[] =\n")
	for i := uint(0); i < mat.g_size*mat.v_size; i++ {
		st = fmt.Sprintf("%.3f\t", mat.elements[i])
		file.WriteString(st + " ")
		if i != 0 && (i+1)%mat.v_size == 0 {
			file.WriteString("\n")
		}
	}
	file.WriteString("\n" + time.Now().Local().String() + "\n")
	file.Close()
}

// Метод виведення матриці на екран
//  - друкує матрицю *Matrix
func (mat *Matrix) PrintMatrix() {
	for i := uint(0); i < mat.g_size*mat.v_size; i++ {
		fmt.Printf("%.3f\t", mat.elements[i])
		if i != 0 && (i+1)%mat.v_size == 0 {
			fmt.Println()
		}
	}
}

// Метод додавання матриць
//  A[] = A[] + B[]
// Вхідні дані:
//  *Matrix - матриця А
//  Matrix oper - матриця B
// Вихідні дані:
//  *Matrix - матриця (А) A[] = А[] + B[]
//  int:
//   -1 - помилка. матриці різної розмірності
//   0 - помилки відсутні
func (mat *Matrix) Add(oper Matrix) int {
	if (mat.g_size != oper.g_size) || (mat.v_size != oper.g_size) {
		return -1
	}
	for i := range mat.elements {
		mat.elements[i] += oper.elements[i]
	}
	return 0
}

// Метод віднімання матриць
//  A[] = A[] - B[]
// Вхідні дані:
//  *Matrix - матриця А
//  Matrix oper - матриця B
// Вихідні дані:
//  *Matrix - матриця (А) A[] = А[] - B[]
//  int:
//   -1 - помилка. матриці різної розмірності
//   0 - помилки відсутні
func (mat *Matrix) Sub(oper Matrix) int {
	if (mat.g_size != oper.g_size) || (mat.v_size != oper.v_size) {
		return -1
	}
	for i := range mat.elements {
		mat.elements[i] -= oper.elements[i]
	}
	return 0
}

// Метод множення матриць
//  A[] = A[] * B[]
// Вхідні дані:
//  *Matrix - матриця А
//  Matrix oper - матриця B
// Вихідні дані:
//  *Matrix - матриця (А) A[] = А[] * B[]
//  int:
//   -1 - помилка. матриці не задовільняють умові множення
//   0 - помилки відсутні
func (mat *Matrix) Multiple(oper Matrix) int {
	if (mat.g_size != oper.v_size) || (mat.v_size != oper.g_size) {
		return -1
	}
	m := 0
	temp := mat.MatrixCopy()
	*mat = setMatrix(temp.g_size, oper.v_size)
	for j := uint(0); j < temp.g_size; j++ {
		for i := uint(0); i < oper.v_size; i++ {
			for k := uint(0); k < temp.v_size; k++ {
				if k == 0 {
					mat.elements[m] = 0
				}
				mat.elements[m] += temp.elements[j*temp.v_size+k] * oper.elements[k*oper.v_size+i]
			}
			m++
		}
	}
	return 0
}

// Метод ділення матриць
//  A[] = A[] / B[]
// Вхідні дані:
//  *Matrix - матриця А
//  Matrix oper - матриця B
// Вихідні дані:
//  *Matrix - матриця (А) A[] = А[] / B[]
//  int:
//   -2 - оберненої матриці не знайдено
//   -1 - помилка. матриці не задовільняють умові ділення
//   0 - помилки відсутні
func (mat *Matrix) Division(oper Matrix) int {
	if oper.RepMatrix() != 0 {
		return -1
	}
	if mat.Multiple(oper) != 0 {
		return -2
	}
	return 0
}

// Метод додавання числа
//  A[] = A[] + float(B)
// Вхідні дані:
//  *Matrix - матриця А
//  float oper - число B
// Вихідні дані:
//  *Matrix - матриця (А) A[] = А[] + float(B)
func (mat *Matrix) VAdd(oper float32) {
	for i := range mat.elements {
		mat.elements[i] += oper
	}
}

// Метод віднімання числа
//  A[] = A[] - float(B)
// Вхідні дані:
//  *Matrix - матриця А
//  float oper - число B
// Вихідні дані:
//  *Matrix - матриця (А) A[] = А[] - float(B)
func (mat *Matrix) VSub(oper float32) {
	for i := range mat.elements {
		mat.elements[i] -= oper
	}
}

// Метод множення числа
//  A[] = A[] * float(B)
// Вхідні дані:
//  *Matrix - матриця А
//  float oper - число B
// Вихідні дані:
//  *Matrix - матриця (А) A[] = А[] * float(B)
func (mat *Matrix) VMultiple(oper float32) {
	for i := range mat.elements {
		mat.elements[i] *= oper
	}
}

// Метод додавання числа
//  A[] = A[] / float(B)
// Вхідні дані:
//  *Matrix - матриця А
//  float oper - число B
// Вихідні дані:
//  *Matrix - матриця (А) A[] = А[] / float(B)
//  int:
//   -1 - ділення на нуль
//   0 - нормальне виконання
func (mat *Matrix) VDivision(oper float32) int {
	if oper == 0 {
		return -1
	}
	for i := range mat.elements {
		mat.elements[i] /= oper
	}
	return 0
}

// Метод поелементного множення матриць
//  A[] = A[] e*l B[]
// Вхідні дані:
//  *Matrix - матриця А
//  Matrix oper - матриця B
// Вихідні дані:
//  *Matrix - матриця (А) A[] = А[] e*l B[]
//  int:
//   -1 - помилка. матриці різної розмірності
//   0 - помилки відсутні
func (mat *Matrix) ElemMultiple(oper Matrix) int {
	if (mat.g_size != oper.g_size) || (mat.v_size != oper.g_size) {
		return -1
	}
	for i := range mat.elements {
		mat.elements[i] *= oper.elements[i]
	}
	return 0
}

// Метод поелементного ділення матриць
//  A[] = A[] e|l B[]
// Вхідні дані:
//  *Matrix - матриця А
//  Matrix oper - матриця B
// Вихідні дані:
//  *Matrix - матриця (А) A[] = А[] e|l B[]
//  int:
//   -1 - помилка. матриці різної розмірності
//   0 - помилки відсутні
//   >0 - номер елементу що викликав дылення на нуль
func (mat *Matrix) ElemDivision(oper Matrix) int {
	if (mat.g_size != oper.g_size) || (mat.v_size != oper.g_size) {
		return -1
	}
	for i := range mat.elements {
		if oper.elements[i] == 0 {
			return i
		}
	}
	for i := range mat.elements {
		mat.elements[i] *= oper.elements[i]
	}
	return 0
}

// Метод знаходження оберненої матриці
//  - обернену матрицю має лише квадратна невироджена матриця
// Вхідні дані:
//  *Matrix - вхідна матриця Matrix
// Вихідня дані:
//  *Matrix - обернена матриця Matrix
//  int:
//   -2 - оберненої матриці не існує
//   -1 - матриця не є квадратною
//   0 - нормальне виконання
func (mat *Matrix) RepMatrix() int {
	if mat.g_size != mat.v_size {
		return -1
	}
	var i, j, n, res int
	n = int(mat.g_size)
	y := make([][]float32, n)
	for i := range y {
		y[i] = make([]float32, n)
	}
	a := mat.toArray()
	b := make([]float32, n)
	x := make([]float32, n)
	for i = 0; i < n; i++ {
		for j = 0; j < n; j++ {
			if j == i {
				b[j] = 1
			} else {
				b[j] = 0
			}
		}
		res = slou(a, n, b, x)
		if res != 0 {
			break
		} else {
			for j = 0; j < n; j++ {
				y[j][i] = x[j]
			}
		}
	}
	*mat = mat.toMatrix(y)
	if res != 0 {
		return -2
	} else {
		return 0
	}
}

// Метод знаходження детермінанту матриці
//  - детермінант можливо знайти лише для квадратної матриці
// Вхідні дані:
//  *Matrix - вхідна матриця Matrix
// Вихідня дані:
//  float, int:
//   0, -1 - матриця не є квадратною
//   det ,0 - нормальне виконання
func (mat *Matrix) DetMatrix() (float32, int) {
	if mat.g_size != mat.v_size {
		return 0, -1
	}
	matrix_a := mat.toArray()
	var i, j, k, r uint
	var c, M, det float32
	var max float64
	det = 1
	a := make([][]float32, mat.v_size)
	for i := range a {
		a[i] = make([]float32, mat.g_size)
	}
	for i = 0; i < mat.g_size; i++ {
		for j = 0; j < mat.g_size; j++ {
			a[i][j] = matrix_a[i][j]
		}
	}
	for k = 0; k < mat.g_size; k++ {
		max = math.Abs(float64(a[k][k]))
		r = k
		for i = k + 1; i < mat.g_size; i++ {
			if math.Abs(float64(a[i][k])) > max {
				max = math.Abs(float64(a[i][k]))
				r = i
			}
		}
		if r != k {
			det = -det
		}
		for j = 0; j < mat.g_size; j++ {
			c = a[k][j]
			a[k][j] = a[r][j]
			a[r][j] = c
		}
		for i = k + 1; i < mat.g_size; i++ {
			M = a[i][k] / a[k][k]
			for j = k; j < mat.g_size; j++ {
				a[i][j] -= M * a[k][j]
			}
		}
	}
	for i = 0; i < mat.g_size; i++ {
		det *= a[i][i]
	}

	return det, 0
}

// (Privat) Функція вирішення СЛАР методом Гаусса
//  - використовуеться для знаходження оберненої матриці
// Вхідні дані:
//  [][]float32 matrica_a - вхідна матриця
//  int n - розмірність системи
//  []float32 mas_b - вектор вільних коефіцієнтів
//  []float32 x - вектор рішень i-ї системи рівнянь (стовпець шуканої матриці)
// Вихідні дані:
//  []float32 x - стовпець шуканої матриці
//  int - коди помилок
func slou(matrica_a [][]float32, n int, mas_b []float32, x []float32) int {
	var i, j, k, r int
	var c, M, s float32
	var max float64
	a := make([][]float32, n)
	for i := range a {
		a[i] = make([]float32, n)
	}
	b := make([]float32, n)
	for i = 0; i < n; i++ {
		for j = 0; j < n; j++ {
			a[i][j] = matrica_a[i][j]
		}
	}
	for i = 0; i < n; i++ {
		b[i] = mas_b[i]
	}
	for k = 0; k < n; k++ {
		max = math.Abs(float64(a[k][k]))
		r = k
		for i = k + 1; i < n; i++ {
			if math.Abs(float64(a[i][k])) > max {
				max = math.Abs(float64(a[i][k]))
				r = i
			}
		}
		for j = 0; j < n; j++ {
			c = a[k][j]
			a[k][j] = a[r][j]
			a[r][j] = c
		}
		c = b[k]
		b[k] = b[r]
		b[r] = c
		for i = k + 1; i < n; i++ {
			M = a[i][k] / a[k][k]
			for j = k; j < n; j++ {
				a[i][j] -= M * a[k][j]
			}
			b[i] -= M * b[k]
		}
	}
	if a[n-1][n-1] == 0 {
		if b[n-1] == 0 {
			return -1
		} else {
			return -2
		}
	} else {
		for i = n - 1; i >= 0; i-- {
			s = 0
			for j = i + 1; j < n; j++ {
				s += a[i][j] * x[j]
			}
			x[i] = (b[i] - s) / a[i][i]
		}
		return 0
	}
}
