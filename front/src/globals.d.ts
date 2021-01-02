/**
 * fix: https://stackoverflow.com/questions/30429520/adding-existing-functions-in-typescript-node-getattribute
 */
interface EventTarget {
    getAttribute(attr: string): string;
    hasAttribute(attr: string): boolean;
}