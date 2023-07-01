#include <stdbool.h>
#include <stdint.h>

void *get_bpe_from_model(const char *model);
uint32_t *encode(void *ptr, const char *prompt);
void free_bpe(void *ptr);
