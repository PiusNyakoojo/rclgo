# rclgo
ROS Client library for Go

This is a Go wrapper for [ROS2]()'s client library. It is designed to provide an interface similar to rclcpp.

## TODO
Inspired by [rclcs](https://github.com/firesurfer/rclcs)

* [ ] Initialize and de-initialize client RCL
    - Handle errors of returned values, throwing exceptions where useful
    - Implement rmw (ROS middleware) error handling methods
* [ ] Create a Node
* [ ] Create a Publisher
* [ ] Create a Subscription
* [ ] Create a Service/Client
    - Get a request and answer it
* [ ] Publish and receive a message
* [ ] Generate code for messages
* [ ] Setting qos profiles
* [ ] Use the currently implemented graph functions
* [ ] Have intellisense comments
