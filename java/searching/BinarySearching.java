import java.util.Scanner;

class BinarySearching {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        try {
            // Input
            System.out.print("Input jumlah data: ");
            int amountOfData = scanner.nextInt();
            System.out.print("Input item data (setiap item dipisahkan dengan spasi): ");
            int[] itemData = new int[amountOfData];
            for (int a = 0; a < amountOfData; a++) {
                itemData[a] = Integer.parseInt(scanner.next());
            }
            System.out.print("Input item data yang dicari: ");
            int itemSearch = scanner.nextInt();

            // Process
            // Sorting
            for (int a = 0; a < amountOfData - 1; a++) {
                int itemA = itemData[a];
                int indexLowest = a + 1;
                int lowest = itemData[indexLowest];
                for (int b = a + 1; b < amountOfData; b++) {
                    int itemB = itemData[b];
                    if (itemB < indexLowest) {
                        indexLowest = b;
                        lowest = itemB;
                    }
                }
                if (lowest < itemA) {
                    itemData[a] = lowest;
                    itemData[indexLowest] = itemA;
                }
            }

            // Binary searching
            int count = 0;
            boolean isFound = false;
            int initialIndex = 0;
            int finalIndex = amountOfData - 1;
            int tempValue = 0;
            while (!isFound) {
                int middleIndex = (initialIndex + finalIndex) / 2;
                int value = itemData[middleIndex];
                if (value == itemSearch) {
                    isFound = true;
                    break;
                } else if (value < itemSearch) {
                    initialIndex = middleIndex + 1;
                } else if (value > itemSearch) {
                    finalIndex = middleIndex - 1;
                }
                if (value == tempValue) {
                    count += 1;
                } else {
                    tempValue = value;
                    count = 0;
                }
                if (count == 3) {
                    break;
                }   
            }
            if (isFound) {
                System.out.println("Output: data yang dicari ditemukan");
            } else {
                System.out.println("Output: data yang dicari tidak ditemukan");
            }
        } catch (Exception e) {
            e.printStackTrace();
            System.out.println("Invalid data");
        }
    }
}
