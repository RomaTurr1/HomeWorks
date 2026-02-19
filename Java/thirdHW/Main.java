package Java.thirdHW;
import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);

        // Ввод массива
        System.out.print("Введите размер массива: ");
        int n = sc.nextInt();

        int[] arr = new int[n];

        System.out.println("Введите элементы массива:");
        for (int i = 0; i < n; i++) {
            arr[i] = sc.nextInt();
        }

        // 1. Вывод в обратном порядке
        System.out.print("Обратный порядок: ");
        for (int i = n - 1; i >= 0; i--) {
            System.out.print(arr[i] + " ");
        }
        System.out.println();

        // 2. Подсчет нулей
        int zeroCount = 0;
        for (int x : arr) {
            if (x == 0) zeroCount++;
        }
        System.out.println("Количество нулей: " + zeroCount);

        // 3. Сумма элементов
        int sum = 0;
        for (int x : arr) {
            sum += x;
        }
        System.out.println("Сумма: " + sum);

        // 4. Максимальный элемент
        int max = arr[0];
        for (int x : arr) {
            if (x > max) max = x;
        }
        System.out.println("Максимум: " + max);

        // 5. Проверка на число 10
        boolean found10 = false;
        for (int x : arr) {
            if (x == 10) {
                found10 = true;
                break;
            }
        }

        if (found10) System.out.println("YES");
        else System.out.println("NO");

        sc.close();
    }
}