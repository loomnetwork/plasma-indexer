pragma solidity >=0.4.21 <0.6.0; 

contract LoomStore {
  uint value;
  string name;

  event NewValueSet(string _name, uint _value);

  function set(string memory _name, uint _value) public {
    value = _value;
    name = _name;
    emit NewValueSet(name, value);
  }

  function getValue() public view returns (uint) {
    return value;
  }

  function getName() public view returns (string memory) {
    return name;
  }
}
