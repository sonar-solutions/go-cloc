// This is a single-line comment
function foo() {
    // This is a block comment
    /*
      This is a block comment that spans multiple lines
      It's used to illustrate the issue with multi-line comments
    */
    console.log("Hello, World!");
  }
  
  /*
  This is a multi-line comment that spans multiple lines
  It should not be counted as a line of code
  
  But wait, what about this?
  */
  
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
  
  /**
   * This is a JSDoc-style comment that spans multiple lines
   * It's used to illustrate the issue with multi-line comments
   */
  function baz() {
    return "hello";
  }
  
  if (true) {
    // This is an if statement with a block comment
    /*
      This is a block comment that spans multiple lines
      It's used to illustrate the issue with multi-line comments
    */
    console.log("Hello, World!");
  }
