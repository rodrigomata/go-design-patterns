# Builder Design Pattern

The Builder pattern is used when the construction of complex objects is needed without instantiating their struct.

## Characteristics

- Removes the creation logic for all the objects
- Can be reused to create many types
- Abstracts complex creations
- Creates objects step by step

## Notes

All methods from the interface must be implemented with this pattern; however, it helps to maintain an unpredictable number of objects, so it should be used when the creation algorithm is steady and has no frequent changes.
