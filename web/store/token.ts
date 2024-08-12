import {defineStore} from "pinia";
import axios from "axios";
import type {TokenState} from "./types";

const baseUrl = "http://localhost:8083";

export const useTokenState = defineStore("token", {
    state: (): TokenState => ({
        token: null,
        error: false,
        loading: true,
    }),
    actions: {
        async fetchToken(code: string) {
            this.loading = true;

            try {
                const {data} = await axios.get<string>(baseUrl + "/token?code=" + code);
                this.token = data;
                localStorage.setItem("token", data)
                this.error = false;
            } catch (error) {
                this.error = true;
            } finally {
                this.loading = false;
            }
        },
        removeToken() {
            this.token = null;
            localStorage.removeItem("token");
            this.error = false;
            this.loading = false;
        }
    }
});