// https://github.com/ros2/rclcpp/blob/master/rclcpp/src/rclcpp/callback_group.cpp

package callback_group

type CallbackGroup struct {
}

// TODO: Return std::vector<rclcpp::subscription::SubscriptionBase::WeakPtr>
func (cbg *CallbackGroup) GetSubscriptionPtrs() {

}

// TODO: Return std::vector<rclcpp::timer::TimerBase::WeakPtr> &

func (cbg *CallbackGroup) GetTimerPtrs() {

}

// TODO: Return const std::vector<rclcpp::service::ServiceBase::WeakPtr> &

func (cbg *CallbackGroup) GetServicePtrs() {

}

// TODO: Return const std::vector<rclcpp::client::ClientBase::WeakPtr> &

func (cbg *CallbackGroup) GetClientPtrs() {

}

// TODO: Return std::atomic_bool &

func (cbg *CallbackGroup) CanBeTakenFrom() {

}

// TODO: Return const CallbackGroupType &

func (cbg *CallbackGroup) Type() {

}

// TODO: Accept const rclcpp::subscription::SubscriptionBase::SharedPtr subscription_ptr

func (cbg *CallbackGroup) AddSubscription() {

}

// TODO: Accept const rclcpp::timer::TimerBase::SharedPtr

func (cbg *CallbackGroup) AddTimer() {

}

// TODO: Accept const rclcpp::service::ServiceBase::SharedPtr service_ptr

func (cbg *CallbackGroup) AddService() {

}

// TODO:

func (cbg *CallbackGroup) AddClient() {

}
