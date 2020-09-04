//     MEMORY 
//
// STACK Base  HIGH
// SP + 24   // This will be return address
// SP + 16   // Argument 2
// SP + 8    // Argument 1
// SP        // Stack Top
//						 LOW
// Stack size is SP - STACK TOP == 24
// In call instructment
// It will keep return address at 0(SP)