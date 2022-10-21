pragma solidity >=0.7.0 <0.9.9;

contract Todo{

    Task[] tasks;

    struct Task {
        string content;
        bool status;
    }

    function add(string memory _content) public {
        tasks.push(Task(_content, false));
    }

    function get(uint _id) public view returns (Task memory){
        return tasks[_id];
    }

    function list() public view returns (Task[] memory){
        return tasks;
    }

    function update(uint _id, string memory _content) public {
        tasks[_id].content = _content;
    }

    function remove(uint _id) public{
        delete tasks[_id];
    }
}
