import * as api from '../api';

export function fetchGoalsSucceeded(goals) {
    return {
        type: 'FETCH_GOALS_SUCCEEDED',
        payload: {
            goals
        }
    }
}

export function fetchGoals() {
    return dispatch => {
        api.fetchGoals().then(resp => {
            dispatch(fetchGoalsSucceeded(resp.data))
        })
    }
}


function createGoalSucceeded(goal) {
    return {
        type: 'CREATE_GOAL_SUCCEEDED',
        payload: {
            goal
        }
    }
}

export function createGoal({goalName, percentComplete}) {
    return dispatch => {
        api.createGoal({goalName, percentComplete}).then(resp => {
            dispatch(createGoalSucceeded(resp.data))
        })
    };
}