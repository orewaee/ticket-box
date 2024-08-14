import {defineStore} from "pinia";
import axios from "axios";
import type {Topic, TopicsState} from "./types";

const baseUrl = "http://localhost:8083";

export const useTopicsState = defineStore("topics", {
    state: (): TopicsState => ({
        topics: null,
        error: false,
        loading: true
    }),
    actions: {
        async fetchTopics(guildId: string) {
            this.loading = true;

            try {
                const {data} = await axios.get<Topic[]>(baseUrl + "/topics/" + guildId, {
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("token")
                    }
                });
                this.topics = data;
                this.error = false;
            } catch (error) {
                this.error = true;
            } finally {
                this.loading = false;
            }
        },
        async createTopic(guildId: string, emoji: string, name: string, description: string) {
            try {
                const {data} = await axios<Topic>({
                    method: "POST",
                    url: baseUrl + "/topic",
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("token")
                    },
                    data: {guild_id: guildId, emoji, name, description}
                })

                this.topics = [...(this.topics != null ? this.topics : []), data];

                this.error = false;
            } catch (error) {

            }
        },
        async removeTopicById(topicId: string) {
            try {
                await axios({
                    method: "DELETE",
                    url: baseUrl + "/topic/" + topicId,
                    headers: {
                        Authorization: "Bearer " + localStorage.getItem("token")
                    }
                })

                this.topics = (this.topics != null ? this.topics : []).filter(topic => topic.id != topicId)

                this.error = false;
            } catch (error) {

            }
        },
    }
});