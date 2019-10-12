import React from 'react';
import {Badge, Col, Progress, Row} from 'reactstrap';

const Task = props => {
    return (
        <div>
            Id: {props.id} <br/>
            UserId: {props.userId} <br/>
            TaskTypeId: {props.taskTypeId} <br/>
            Title: {props.title} <br/>
            UnitId: {props.unitId} <br/>
            Size: {props.size} <br/>
            Default Tags: {props.defaultTags} <br/>
            <br/>
        {/*<Row>*/}
        {/*    <Col xs="4">{props.title} <Badge color="danger">{props.taskTypeId}</Badge></Col>*/}
        {/*    <Col>{props.unitId}{props.size}</Col>*/}
        {/*    <Col><Progress value={props.percentComplete}>{props.percentComplete}%</Progress></Col>*/}
        {/*</Row>*/}
        </div>
    );
};

// Goal.propTypes = {
//     goalName: PropTypes.string.isRequired,
//     percentComplete: PropTypes.number.isRequired,
// };

export default Task;
