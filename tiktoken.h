#include <stdbool.h>
#include <stdint.h>

// 1-to-1 mapping to the functions in lib.rs
void *get_bpe_from_model(const char *model);
void free_bpe(void *ptr);
uint32_t *encode(void *ptr, const char *prompt, uint32_t *len /*output length*/);
