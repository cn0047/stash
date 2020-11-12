#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <assert.h>

struct word {
    char* content;
};

typedef struct word Word;

struct sentence {
    struct word* content;
    int words_count;
};

typedef struct sentence Sentence;

struct sentence* get_sentence()
{
    int words_count = 2;

    Word *w1 = malloc(1 * sizeof(Word*));
    w1->content = "it";
    Word *w2 = malloc(1 * sizeof(Word*));
    w2->content = "works";

    Sentence *s = malloc(1 * sizeof(Sentence*));
    s->words_count = words_count;
    s->content = malloc(words_count * sizeof(Word));
    s->content[0] = *w1;
    s->content[1] = *w2;

    return s;
}

void print_sentence(struct sentence* s)
{
    for (int i = 0; i < s->words_count; i++) {
        printf("%s ", s->content[i].content);
    }
    printf("\n");
}

int main()
{
    Sentence *s = get_sentence();
    print_sentence(s);

    return 0;
}
