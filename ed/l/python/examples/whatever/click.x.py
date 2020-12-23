import click


@click.command()
@click.option('--compile', is_flag=True, default=False, help='Compile ')
def main(compile):
  print(compile)


if __name__ == '__main__':
    main()
