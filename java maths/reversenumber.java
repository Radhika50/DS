
import java.util.Scanner;

public class reversenumber {
  public static void main(String[] args) {
      int number;
      Scanner sc=new Scanner(System.in);
      number=sc.nextInt();
      sc.close();
      int rev=0;
      while(number > 0){
        int rem=number%10;
        rev=(rev*10)+rem;
        number=number/10;
      }
      System.out.println("Reeverse "+rev);
  }
}
