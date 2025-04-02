import java.util.Scanner;

public class countdigits{
  public static void main(String[] args){
    Scanner sc=new Scanner(System.in);
    System.out.print("Enter number: ");
    int n=sc.nextInt();
    sc.close();
    int cnt=0;
    while(n>0){
      n=n/10;
      cnt++;
    }
    System.out.println(cnt);
  } 
}