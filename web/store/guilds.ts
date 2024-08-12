import {defineStore} from "pinia";
import axios from "axios";
import type {Guild, GuildsState} from "./types";

const baseUrl = "http://localhost:8083";

export const useGuildsState = defineStore("guilds", {
    state: (): GuildsState => ({
        guilds: null,
        error: false,
        loading: true
    }),
    actions: {
        async fetchGuilds() {
            this.loading = true;

            try {
                const {data} = await axios.get<Guild[]>(baseUrl + "/guilds", {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("token")
                    }
                });
                this.guilds = data;
                this.error = false;
            } catch (error) {
                this.error = true;
            } finally {
                this.loading = false;
            }
        }
    }
});