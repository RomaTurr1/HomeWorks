package Java.fourthHW;

public class Main {

    public static void main(String[] args) {

        multiplicationTable();
        rectangle();
        triangle();
        countEvenNested();
        sumNested();

        int[][] matrix = createMatrix();
        printMatrix(matrix);
        sumMatrix(matrix);
        maxMatrix(matrix);
        countEvenMatrix(matrix);
    }

    // 1. Таблица умножения 10x10
    static void multiplicationTable() {
        System.out.println("Таблица умножения:");
        for (int i = 1; i <= 10; i++) {
            for (int j = 1; j <= 10; j++) {
                System.out.print(i * j + "\t");
            }
            System.out.println();
        }
        System.out.println();
    }

    // 2. Прямоугольник n x m
    static void rectangle() {
        int n = 4; // строки
        int m = 6; // столбцы

        System.out.println("Прямоугольник:");
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < m; j++) {
                System.out.print("* ");
            }
            System.out.println();
        }
        System.out.println();
    }

    // 3. Треугольник
    static void triangle() {
        System.out.println("Треугольник:");
        for (int i = 1; i <= 5; i++) {
            for (int j = 0; j < i; j++) {
                System.out.print("*");
            }
            System.out.println();
        }
        System.out.println();
    }

    // 4. Чётные от 1 до 100 (вложенные циклы)
    static void countEvenNested() {
        int count = 0;

        for (int tens = 0; tens < 10; tens++) {
            for (int i = 1; i <= 10; i++) {
                int num = tens * 10 + i;
                if (num % 2 == 0) count++;
            }
        }

        System.out.println("Чётных от 1 до 100: " + count);
        System.out.println();
    }

    // 5. Сумма от 1 до 100 (вложенные циклы)
    static void sumNested() {
        int sum = 0;

        for (int tens = 0; tens < 10; tens++) {
            for (int i = 1; i <= 10; i++) {
                int num = tens * 10 + i;
                sum += num;
            }
        }

        System.out.println("Сумма от 1 до 100: " + sum);
        System.out.println();
    }

    // 6. Создание 3x3 массива
    static int[][] createMatrix() {
        int[][] m = {
                {1, 2, 3},
                {4, 5, 6},
                {7, 8, 9}
        };
        return m;
    }

    // Вывод матрицы
    static void printMatrix(int[][] m) {
        System.out.println("Матрица:");
        for (int i = 0; i < m.length; i++) {
            for (int j = 0; j < m[i].length; j++) {
                System.out.print(m[i][j] + "\t");
            }
            System.out.println();
        }
        System.out.println();
    }

    // 7. Сумма элементов матрицы
    static void sumMatrix(int[][] m) {
        int sum = 0;

        for (int i = 0; i < m.length; i++) {
            for (int j = 0; j < m[i].length; j++) {
                sum += m[i][j];
            }
        }

        System.out.println("Сумма матрицы: " + sum);
        System.out.println();
    }

    // 8. Максимум в матрице
    static void maxMatrix(int[][] m) {
        int max = m[0][0];

        for (int i = 0; i < m.length; i++) {
            for (int j = 0; j < m[i].length; j++) {
                if (m[i][j] > max) max = m[i][j];
            }
        }

        System.out.println("Максимум в матрице: " + max);
        System.out.println();
    }

    // 9. Количество чётных
    static void countEvenMatrix(int[][] m) {
        int count = 0;

        for (int i = 0; i < m.length; i++) {
            for (int j = 0; j < m[i].length; j++) {
                if (m[i][j] % 2 == 0) count++;
            }
        }

        System.out.println("Чётных в матрице: " + count);
        System.out.println();
    }
}