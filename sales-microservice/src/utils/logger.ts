import { createLogger, format, transports, Logger } from 'winston';

export const logger: Logger = createLogger({
  level: 'info',
  format: format.combine(
    format.timestamp(),
    format.printf(({ level, message, timestamp }) => {
      return `${timestamp} [${level}]: ${message}`;
    })
  ),
  transports: [new transports.Console()],
});
