import { validate } from 'class-validator';

export async function validateClass<T extends object>(instance: T): Promise<void> {
    const errors = await validate(instance);
    if (errors.length > 0) {
        throw new Error('Validation failed');
    }
}