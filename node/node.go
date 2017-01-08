// https://github.com/ros2/rclcpp/blob/master/rclcpp/src/rclcpp/node.cpp

package node

import "time"

type Node struct {
	name string
	// rclcpp::context::Context::SharedPtr context,
	base       Base
	graph      Graph
	timers     Timers
	topics     Topics
	services   Services
	parameters Parameters
}

// TODO: Accept values for base graph, etc..?
func NewNode() *Node {
	var n Node

	return &n
}

// TODO: Write test
func (n *Node) GetName() string {
	return ""
}

// TODO: Accept rclcpp::callback_group::CallbackGroupType group_type
// TODO: Return rclcpp::callback_group::CallbackGroup::SharedPtr
func (n *Node) CreateCallbackGroup(group_type string) {

}

// TODO: Write test
// TODO: Accept rclcpp::callback_group::CallbackGroup::SharedPtr
func (n *Node) GroupInNode(group string) bool {
	return false
}

// TODO: Accept std::vector<rclcpp::parameter::ParameterVariant> & parameters
// TODO: Return std::vector<rcl_interfaces::msg:SetParametersResult
func (n *Node) SetParameters(parameters []string) {

}

// TODO: Accept std::vector<rclcpp::parameter::ParameterVariant> & parameters
// TODO: Return rcl_interfaces::msg::SetParametersResult
func (n *Node) SetParametersAtomically(parameters []string) {

}

// TODO: Accept std::vector<std::string> & names
// TODO: Return std::vector<rclcpp::parameter::ParameterVariant>
func (n *Node) GetParameters(names []string) {

}

// TODO: Accept const std::string & name
// TODO: Return rclcpp::parameter::ParameterVariant
func (n *Node) GetParameter(name string) {

}

// TODO: Write tests
// TODO: Accept const std::string & name, rclcpp::parameter::ParameterVariant & parameter
func (n *Node) HasParameter(name, parameter string) bool {
	return false
}

// TODO: Accept std::vector<std::string> & names
// TODO: Return std::vector<rcl_interfaces::msg::ParameterDescriptor
func (n *Node) DescribeParameters(names []string) {

}

// TODO: Accept std::vector<std::string> & names
// TODO: Return std::vector<uint8_t>
func (n *Node) GetParameterTypes(names []string) {

}

// TODO: Accept std::vector<std::string> & prefixes, uint64_t depth
// TODO: Return rcl_interfaces::msg::ListParametersResult
func (n *Node) ListParameters(prefixes []string, depth int) {

}

// TODO: Return std::map<std::string, std::string>
func (n *Node) GetTopicNamesAndTypes() {

}

// TODO: Accept size_t
// TODO: Return const std::string & topic_name
func (n *Node) CountPublishers(topic_name string) {

}

// TODO: Accept size_t
// TODO: Return const std::string * topic_name
func (n *Node) CountSubscribers(topic_name string) {

}

// TODO: Return std::vector<rclcpp::callback_group::CallbackGroup::WeakPtr> &
func (n *Node) GetCallbackGroups() {

}

// TODO: Return rclcpp::event::Event::SharedPtr
func (n *Node) GetGraphEvent() {

}

// TODO: Accept rclcpp::event::Event::SharedPtr event, std::chrono::nanoseconds timeout
// TODO: Return void
func (n *Node) WaitForGraphChange(event string, timeout time.Time) {

}

// TODO: Return rclcpp::node_interfaces::NodeBaseInterface::SharedPtr
func (n *Node) GetBaseInterface() {

}

// TODO: Return rclcpp::node_interfaces::NodeGraphInterface::SharedPtr
func (n *Node) GetGraphInterface() {

}

// TODO: Return rclcpp::node_interfaces::NodeTimersInterface::SharedPtr
func (n *Node) GetTimersInterface() {

}

// TODO: Return rclcpp::node_interfaces::NodeTopicsInterface::SharedPtr
func (n *Node) GetTopicsInterface() {

}

// TODO: Return rclcpp::node_interfaces::NodeServicesInterface::SharedPtr
func (n *Node) GetServicesInterface() {

}

// TODO: Return rclcpp::node_interfaces::NodeParametersInterface::SharedPtr
func (n *Node) GetParametersInterface() {

}
