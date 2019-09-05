CC=clang

.PHONY: clean

hello: hello.o main.o
	$(CC) -o hello hello.o main.o

hello.o: hello.c
	$(CC) -c hello.c

main.o: main.c
	$(CC) -c main.c

libhello.a: hello.o
	ar rusv libhello.a hello.o

clean:
	rm -rf *.o lib*.a hello
