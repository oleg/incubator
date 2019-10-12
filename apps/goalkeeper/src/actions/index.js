import * as api from '../api';

export function fetchTasksSucceeded(tasks) {
    return {
        type: 'FETCH_TASKS_SUCCEEDED',
        payload: {
            tasks
        }
    }
}

export function fetchTasks() {
    return dispatch => {
        api.fetchTasks().then(resp => {
            dispatch(fetchTasksSucceeded(resp.data))
        })
    }
}


function createTaskSucceeded(task) {
    return {
        type: 'CREATE_TASK_SUCCEEDED',
        payload: {
            task
        }
    }
}

export function createTask({taskName, percentComplete}) {
    return dispatch => {
        api.createTask({taskName, percentComplete}).then(resp => {
            dispatch(createTaskSucceeded(resp.data))
        })
    };
}