import {defineStore} from "pinia";
import type {Guild, GuildState} from "./types";
import axios from "axios";

const baseUrl = "http://localhost:8083";

export const useGuildState = defineStore("guild", {
    state: (): GuildState => ({
        guild: null,
        error: false,
        loading: true
    }),
    getters: {
        id: (state) => state.guild?.id,
        name: (state) => state.guild?.name,
        iconUrl: (state) => state.guild?.icon_url,
        withBot: (state) => state.guild?.with_bot
    },
    actions: {
        async fetchGuild(guildId: string) {
            this.loading = true;

            try {
                const {data} = await axios.get<Guild>(baseUrl + "/guild/" + guildId, {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("token")
                    }
                });
                this.guild = data;
                this.error = false;
            } catch (error) {
                this.error = true;
            } finally {
                this.loading = false;
            }
        }
    }
});