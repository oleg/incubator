export default function goals(state = {goals: []}, action) {
    switch (action.type) {
        case 'FETCH_GOALS_SUCCEEDED': {
            return {
                goals: action.payload.goals
            };
        }
        case 'CREATE_GOAL_SUCCEEDED' : {
            return {
                goals: state.goals.concat(action.payload.goal),
            };
        }
        default: {
            return state;
        }
    }
}