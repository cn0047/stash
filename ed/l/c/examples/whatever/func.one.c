#include <stdio.h>

void f1()
{
    printf("[f1] v1: %s \n", __func__);
}

void print_int_val(int v)
{
    printf("v=%d\n", v);
}

int int_mul_2(int v)
{
    return v * 2;
}

void func_as_param(void (*f) (int))
{
    f(9);
}

void func_as_param_2(int (*f) (int))
{
    int v = f(5);
    print_int_val(v);
}

int main()
{
    // f1();
    // print_int_val(7);
    // func_as_param(print_int_val);
    func_as_param_2(int_mul_2);

    return 0;
}
