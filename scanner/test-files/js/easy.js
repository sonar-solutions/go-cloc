// This is a single-line comment
function foo() {
    // This is a block comment
   
    console.log("Hello, World!");
  }
  
  // This is another single-line comment
  var bar = {
    // This is an object property with a comment
    foo: "hello",
    // This is another object property with a comment
    baz: "world"
  };
  
  (function() {
    // This is an immediately invoked function expression (IIFE)
    // It should not be counted as a line of code
    console.log("Hello, World!");
  })();
  
  // This is the end of the file
  
 
  function baz() {
    return "hello";
  }
  
  if (true) {
    // This is an if statement with a block comment
    console.log("Hello, World!");
  }