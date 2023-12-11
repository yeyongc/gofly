import cookies from "js-cookie"

const tools = {
    Router: {

    },
    Store: {

    },
    LocalStorage: {
        steItem(key:string, value:any){
            localStorage.setItem(key, JSON.stringify(value))
        },
        getItem<T>(key: string):T|string{
            const val = localStorage.getItem(key)
            try {
                return JSON.parse(val as string) as T
            }catch(e){
                return val as string
            }
        }
    },
    Cookie:{
        setItem(key:string, val:any, expire:any){
            cookies.set(key, JSON.stringify(val), expire)
        },
        getItem<T>(key: string):T|string{
            const val = cookies.get(key) as string
            try {
                return JSON.parse(val) as T
            }catch(e){
                return val
            }
        }
    },
    Time: {

    },
    Dom:{

    }
}

export default tools