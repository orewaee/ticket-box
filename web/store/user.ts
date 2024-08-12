import {defineStore} from "pinia";
import axios from "axios";
import type {User, UserState} from "./types";

const baseUrl = "http://localhost:8083";

export const useUserState = defineStore("auth", {
    state: (): UserState => ({
        user: null,
        error: false,
        loading: true,
    }),
    getters: {
        id: (state) => state.user?.id,
        username: (state) => state.user?.username,
        avatarUrl: (state) => state.user?.avatar_url,
    },
    actions: {
        async fetchUser() {
            this.loading = true;

            try {
                const {data} = await axios.get<User>(baseUrl + "/user", {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("token")
                    }
                });
                this.user = data;
                this.error = false;
            } catch (error) {
                this.error = true;
            } finally {
                this.loading = false;
            }
        },
        removeUser() {
            this.user = null;
            this.error = false;
            this.loading = false;
        }
    }
});