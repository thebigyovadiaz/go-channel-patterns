# Wait for Result Pattern
This is a foundational pattern used by larger patterns like fan out/in,
 a goroutine is created to perform some known work and signals their
 result back to the goroutine that created them.
