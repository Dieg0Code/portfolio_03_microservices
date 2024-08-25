import app from "./src/app";
import { logger } from "./src/utils/logger";

const PORT = 8082;

app.listen(PORT, () => {
    logger.info(`Server running on port ${PORT}`);
})