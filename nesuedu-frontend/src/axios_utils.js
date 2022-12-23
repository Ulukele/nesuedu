const API_BASE_URL = "http://localhost:8083/api/v1/"
const API_A_BASE_URL = API_BASE_URL + "auth/"
const API_C_BASE_URL = API_BASE_URL + "content/"

import { createReturnStatement } from '@vue/compiler-core'
import axios from 'axios'

function defaultHeaders(sessionID) {
    return  { 'Content-Type': 'application/json', 'Bearer': sessionID }
}

function defaultConfig(sessionID) {
    return { headers: defaultHeaders(sessionID) }
}

export const axios_utils = {

    tryToSignUp(data_) {
        return axios({
            method: 'post',
            url: API_A_BASE_URL + 'sign-up/',
            data: {
                username: data_.username,
                password: data_.password,
                firstName: data_.firstName,
                lastName: data_.lastName,
            }
        })
    },

    tryToSignIn(data_) {
        return axios({
            method: 'post',
            url: API_A_BASE_URL + 'sign-in/',
            data: {
                username: data_.username,
                password: data_.password
            }
        })
    },

    refreshToken(refreshToken_) {
        return axios({
            method: 'post',
            url: API_A_BASE_URL + 'refresh/',
            data: {
                refreshToken: refreshToken_,
            },
            headers: {'Content-Type': 'application/json'}
        })
    },

    getUser(userId, sessionID) {
        return axios.get(API_C_BASE_URL + 'user/' + userId + '/', defaultConfig(sessionID));
    },

    getTasks(sessionID) {
        return axios.get(API_C_BASE_URL + 'task/', defaultConfig(sessionID));
    },

    createTask(task, sessionID) {
        return axios({
            method: 'post',
            url: API_C_BASE_URL + 'task/',
            data: task,
            headers: defaultHeaders(sessionID)
        });
    },

    deleteTask(taskId, sessionID) {
        return axios.delete(API_C_BASE_URL + 'task/' + taskId + '/', defaultConfig(sessionID));
    },

    createSub(taskId, sessionID) {
        return axios({
            method: 'post',
            url: API_C_BASE_URL + 'task/' + taskId + '/subscribe/',
            headers: defaultHeaders(sessionID)
        });
    },

    deleteSub(taskId, sessionID) {
        return axios({
            method: 'delete',
            url: API_C_BASE_URL + 'task/' + taskId + '/subscribe/',
            headers: defaultHeaders(sessionID)
        });
    },

    getSubs(taskId, sessionID) {
        return axios.get(API_C_BASE_URL + 'task/' + taskId + '/subs/', defaultConfig(sessionID));
    },

}