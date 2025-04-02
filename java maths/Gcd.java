import java.util.Scanner;

public class Gcd {
  public static void main(String[] args) {
      int n1,n2;
      Scanner sc=new Scanner(System.in);
      n1=sc.nextInt();
      n2=sc.nextInt();
      if (n1==0) {System.out.println(n2); return;}
      if (n2==0) {System.out.println(n1); return;}
      //Using Euclidean
      while(n1 > 0 && n2 >0){
        if(n1 > n2) n1=n1%n2;
        else  n2=n2%n1;
      }
      if (n1==0) {
        System.out.println(n2);
      }
      else {
        System.out.println(n1);
      }
  }
}
