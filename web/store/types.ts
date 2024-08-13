export interface TokenState {
    token: string | null
    error: boolean
    loading: boolean
}

export interface User {
    id: string
    username: string
    avatar_url: string
}

export interface UserState {
    user: User | null
    error: boolean
    loading: boolean
}

export interface Guild {
    id: string
    name: string
    icon_url: string
    with_bot: boolean
}

export interface GuildState {
    guild: Guild | null
    error: boolean
    loading: boolean
}

export interface GuildsState {
    guilds: Guild[] | null
    error: boolean
    loading: boolean
}

export interface Topic {
    id: string
    guild_id: string
    emoji: string
    name: string
    description: string
}

export interface TopicsState {
    topics: Topic[] | null
    error: boolean
    loading: boolean
}
