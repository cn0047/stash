zval
-

````
typedef struct _zval_struct {
    zvalue_value value;        /* variable value */
    zend_uint refcount__gc;    /* reference counter */
    zend_uchar type;           /* value type */
    zend_uchar is_ref__gc;     /* reference flag */
} zval;
typedef union _zvalue_value {
    long lval;                 /* long value */
    double dval;               /* double value */
    struct {
        char *val;
        int len;               /* this will always be set for strings */
    } str;                     /* string (always has length) */
    HashTable *ht;             /* an array */
    zend_object_value obj;     /* stores an object store handle, and handlers */
} zvalue_value;
````
