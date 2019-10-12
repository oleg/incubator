export default function tasks(state = {tasks: []}, action) {
    switch (action.type) {
        case 'FETCH_TASKS_SUCCEEDED': {
            return {
                tasks: action.payload.tasks
            };
        }
        case 'CREATE_TASK_SUCCEEDED' : {
            return {
                tasks: state.tasks.concat(action.payload.task),
            };
        }
        default: {
            return state;
        }
    }
}