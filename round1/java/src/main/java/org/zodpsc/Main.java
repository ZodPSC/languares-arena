package org.zodpsc;

public class Main {
    public static void main(String[] args) {
        final int countNumbers = 100000000;
        long begin = System.nanoTime();
        int[] numbers = new int[countNumbers];
        for (int i = 0; i < numbers.length; i++) {
            numbers[i] = (int) (Math.random() * 100);
        }

        long resultTime = System.nanoTime() - begin;

        System.out.println("End generate in time: "+resultTime); //End generate in time: 1 989 613 200
        System.out.println("First and last numbers = "+numbers[0]+" "+numbers[numbers.length - 1]);
    }
}