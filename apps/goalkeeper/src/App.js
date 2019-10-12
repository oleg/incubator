import React, {Component} from 'react';
import {connect} from 'react-redux';
import {Container, Nav, NavItem, NavLink} from 'reactstrap';
import TasksPage from './components/TasksPage'
import {createTask, fetchTasks} from './actions';

class App extends Component {
    componentDidMount() {
        this.props.dispatch(fetchTasks())
    }

    onCreateTask = ({taskName, percentComplete}) => {
        this.props.dispatch(createTask({taskName, percentComplete}));
    }

    render() {
        return (
            <Container>
                <Nav>
                    <NavItem><NavLink href="#">Home</NavLink></NavItem>
                </Nav>
                <TasksPage tasks={this.props.tasks} onCreateTask={this.onCreateTask}/>
            </Container>
        );
    }
}

function mapStateToProps(state) {
    return {
        tasks: state.tasks
    }
}

export default connect(mapStateToProps)(App);
