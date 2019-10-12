import React, {Component} from 'react';
import {Button, Col, Collapse, DropdownItem, DropdownMenu, DropdownToggle, Row, UncontrolledDropdown} from 'reactstrap';
import Task from "./Task"

class TasksPage extends Component {
    constructor(props) {
        super(props)
        this.state = TasksPage.newEmptyForm()
    }

    static newEmptyForm() {
        return {
            showNewGoalForm: false,
            orderBy: 'Name',
            direction: 'Up',
            goalName: '',
            percentComplete: 0
        }
    }

    onGoalNameChange = (e) => {
        this.setState({goalName: e.target.value})
    }

    onPercentCompleteChange = (e) => {
        this.setState({percentComplete: parseInt(e.target.value)});
    }

    resetForm() {
        this.setState(TasksPage.newEmptyForm());
    }

    onCreateTask = (e) => {
        e.preventDefault();
        this.props.onCreateTask({
            goalName: this.state.goalName,
            percentComplete: this.state.percentComplete,
        })
        this.resetForm()
    }

    toggleForm = () => {
        this.setState({showNewGoalForm: !this.state.showNewGoalForm})
    }

    changeOrderBy = (e) => {
        const reverse = () => this.state.direction === 'Up' ? 'Down' : 'Up';
        this.setState({
            orderBy: e.target.textContent,
            direction: this.state.orderBy === e.target.textContent ? reverse() : 'Down'
        })
    }

    render() {
        const byName = function (a, b) {
            return a.title.toUpperCase().localeCompare(b.title.toUpperCase())
        }
        const byPercent = function (a, b) {
            return a.percentComplete - b.percentComplete;
        }
        const byEpic = function (a, b) {
            return a.epic.toUpperCase().localeCompare(b.epic.toUpperCase())
        }
        const byStarted = function (a, b) {
            return a.dateStarted.localeCompare(b.dateStarted)
        }
        var sortFunc;
        switch (this.state.orderBy) {
            case "Name": {sortFunc = byName; break}
            case "Percent Complete": {sortFunc = byPercent; break}
            case "Epic": {sortFunc = byEpic; break}
            case "Started": {sortFunc = byStarted; break}
            default: sortFunc = byName;
        }

        return (
            <div>

                <UncontrolledDropdown>
                    <DropdownToggle size="sm" caret>Sort by: {this.state.orderBy}</DropdownToggle>
                    <DropdownMenu>
                        <DropdownItem onClick={this.changeOrderBy}>Name</DropdownItem>
                        <DropdownItem onClick={this.changeOrderBy}>Percent Complete</DropdownItem>
                        <DropdownItem onClick={this.changeOrderBy}>Epic</DropdownItem>
                        <DropdownItem onClick={this.changeOrderBy}>Started</DropdownItem>
                    </DropdownMenu>
                </UncontrolledDropdown>

                {
                    this.props.tasks
                    .sort((a, b) => this.state.direction === 'Down' ? sortFunc(a, b): sortFunc(b,a))
                    .map(g => <Task key={g.title} {...g} />)
                }
                <Row>
                    <Col>
                        <Button color="primary" onClick={this.toggleForm}>Add Task</Button>
                        <Collapse isOpen={this.state.showNewGoalForm}>
                            <form className="AddGoal" onSubmit={this.onCreateTask}>
                                <input type="text" onChange={this.onGoalNameChange} value={this.state.goalName} placeholder="Goal Name"/>
                                <input type="text" onChange={this.onPercentCompleteChange} value={this.state.percentComplete} placeholder="Percent Complete"/>
                                <button type="submit">Add</button>
                            </form>
                        </Collapse>
                    </Col>
                </Row>

            </div>
        );
    }
}

// GoalsPage.propTypes = {
//     goalName: PropTypes.string.isRequired,
//     percentComplete: PropTypes.number.isRequired,
// };

export default TasksPage;
