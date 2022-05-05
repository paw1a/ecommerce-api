import { useState, useEffect } from 'react';
import axios from "axios";
import {getToken} from "./auth";

const HOST = "localhost";

// axiosClient.defaults.baseURL = 'http://' + HOST + ':8080/api/v1';
//
// axiosClient.defaults.headers = {
//     'Content-Type': 'application/json',
//     'Access-Control-Allow-Origin': '*',
//     'Accept': 'application/json'
// };
//
// axiosClient.defaults.timeout = 6000;

//axiosClient.defaults.withCredentials = true;

// axios.interceptors.response.use(function (response) {
//     return response;
// }, function (error) {
//     if(error.response.status === 401) {
//         console.log("Error 401");
//     }
//     return Promise.reject(error);
// });

export const useAxios = (url, method='get', payload='') => {
    const [data, setData] = useState(null);
    const [loaded, setLoaded] = useState(false);
    const [error, setError] = useState(false);

    useEffect(() => {
        const accessToken = getToken();

        axios({
            method: method,
            url: '/api/v1' + url,
            withCredentials: true,
            data: payload,
            headers: {
                Authorization: 'Bearer ' + accessToken
            }
        })
            .then(resp => {
                const data = resp.data.data;
                console.log(data);
                setData(data);
            })
            .catch(error => {
                console.log(error.response.data);
                setError(error.response.data);
                if (error.response.status === 401) {
                    window.location = '/sign-in';
                }
            })
            .finally(() => setLoaded(true));

    }, [setData, setError, setLoaded]);

    return [data, loaded, error];
}

export const useAuth = (url, method='get', payload='') => {
    const [data, setData] = useState(null);
    const [loaded, setLoaded] = useState(false);
    const [error, setError] = useState(false);

    useEffect(() => {


        axios({
            method: method,
            url: '/api/v1' + url,
            withCredentials: true,
            data: payload,
        })
            .then(resp => {
                const data = resp.data.data;
                console.log(data);
                setData(data);
            })
            .catch(error => {
                console.log(error.response.data);
                setError(error.response.data);
            })
            .finally(() => setLoaded(true));

    }, [setData, setError, setLoaded]);

    return [data, loaded, error];
}

export const apiRequest = (url, method='get', payload='') => {
    let data, loaded, error

    axios({
        method: method,
        url: '/api/v1' + url,
        withCredentials: true,
        data: payload,
    })
        .then(resp => {
            data = resp.data.data;
            console.log(data);
        })
        .catch(err => {
            error = err.response.data
            console.log(error);
        })
        .finally(() => loaded = true);

    return [data, loaded, error]
}
