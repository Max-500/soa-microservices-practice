import express, {Application} from "express";
import { router } from "./Infraestructure/Router/OrderRouter";

const app:Application = express()

app.use(express.json())
app.use("/", router)

const PORT:number = 3001

app.listen(PORT, () => {
    console.log(`SERVER RUNNING IN http://localhost:${PORT}.`);
})