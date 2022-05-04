import { useState, useEffect } from 'react';
import axios from "axios";

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

const useAxios = (url) => {
    const [data, setData] = useState(null);
    const [loaded, setLoaded] = useState(false);
    const [error, setError] = useState(false);

    useEffect(() => {
        axios.get("/api/v1" + url, {withCredentials: true})
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

export default useAxios;