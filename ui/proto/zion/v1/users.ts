/**
 * Generated by the protoc-gen-ts.  DO NOT EDIT!
 * compiler version: 3.19.4
 * source: proto/zion/v1/users.proto
 * git: https://github.com/thesayyn/protoc-gen-ts */
import * as dependency_1 from "./../../protoc-gen-gorm/options/gorm";
import * as pb_1 from "google-protobuf";
export namespace proto.zion.v1 {
    export class User extends pb_1.Message {
        constructor(data?: any[] | {
            id?: number;
            did?: string;
            username?: string;
            name?: string;
            email?: string;
            bio?: string;
            img?: string;
            price_to_message?: number;
            created?: number;
            updated?: number;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], []);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("id" in data && data.id != undefined) {
                    this.id = data.id;
                }
                if ("did" in data && data.did != undefined) {
                    this.did = data.did;
                }
                if ("username" in data && data.username != undefined) {
                    this.username = data.username;
                }
                if ("name" in data && data.name != undefined) {
                    this.name = data.name;
                }
                if ("email" in data && data.email != undefined) {
                    this.email = data.email;
                }
                if ("bio" in data && data.bio != undefined) {
                    this.bio = data.bio;
                }
                if ("img" in data && data.img != undefined) {
                    this.img = data.img;
                }
                if ("price_to_message" in data && data.price_to_message != undefined) {
                    this.price_to_message = data.price_to_message;
                }
                if ("created" in data && data.created != undefined) {
                    this.created = data.created;
                }
                if ("updated" in data && data.updated != undefined) {
                    this.updated = data.updated;
                }
            }
        }
        get id() {
            return pb_1.Message.getField(this, 1) as number;
        }
        set id(value: number) {
            pb_1.Message.setField(this, 1, value);
        }
        get did() {
            return pb_1.Message.getField(this, 2) as string;
        }
        set did(value: string) {
            pb_1.Message.setField(this, 2, value);
        }
        get username() {
            return pb_1.Message.getField(this, 3) as string;
        }
        set username(value: string) {
            pb_1.Message.setField(this, 3, value);
        }
        get name() {
            return pb_1.Message.getField(this, 4) as string;
        }
        set name(value: string) {
            pb_1.Message.setField(this, 4, value);
        }
        get email() {
            return pb_1.Message.getField(this, 5) as string;
        }
        set email(value: string) {
            pb_1.Message.setField(this, 5, value);
        }
        get bio() {
            return pb_1.Message.getField(this, 6) as string;
        }
        set bio(value: string) {
            pb_1.Message.setField(this, 6, value);
        }
        get img() {
            return pb_1.Message.getField(this, 7) as string;
        }
        set img(value: string) {
            pb_1.Message.setField(this, 7, value);
        }
        get price_to_message() {
            return pb_1.Message.getField(this, 8) as number;
        }
        set price_to_message(value: number) {
            pb_1.Message.setField(this, 8, value);
        }
        get created() {
            return pb_1.Message.getField(this, 9) as number;
        }
        set created(value: number) {
            pb_1.Message.setField(this, 9, value);
        }
        get updated() {
            return pb_1.Message.getField(this, 10) as number;
        }
        set updated(value: number) {
            pb_1.Message.setField(this, 10, value);
        }
        static fromObject(data: {
            id?: number;
            did?: string;
            username?: string;
            name?: string;
            email?: string;
            bio?: string;
            img?: string;
            price_to_message?: number;
            created?: number;
            updated?: number;
        }) {
            const message = new User({});
            if (data.id != null) {
                message.id = data.id;
            }
            if (data.did != null) {
                message.did = data.did;
            }
            if (data.username != null) {
                message.username = data.username;
            }
            if (data.name != null) {
                message.name = data.name;
            }
            if (data.email != null) {
                message.email = data.email;
            }
            if (data.bio != null) {
                message.bio = data.bio;
            }
            if (data.img != null) {
                message.img = data.img;
            }
            if (data.price_to_message != null) {
                message.price_to_message = data.price_to_message;
            }
            if (data.created != null) {
                message.created = data.created;
            }
            if (data.updated != null) {
                message.updated = data.updated;
            }
            return message;
        }
        toObject() {
            const data: {
                id?: number;
                did?: string;
                username?: string;
                name?: string;
                email?: string;
                bio?: string;
                img?: string;
                price_to_message?: number;
                created?: number;
                updated?: number;
            } = {};
            if (this.id != null) {
                data.id = this.id;
            }
            if (this.did != null) {
                data.did = this.did;
            }
            if (this.username != null) {
                data.username = this.username;
            }
            if (this.name != null) {
                data.name = this.name;
            }
            if (this.email != null) {
                data.email = this.email;
            }
            if (this.bio != null) {
                data.bio = this.bio;
            }
            if (this.img != null) {
                data.img = this.img;
            }
            if (this.price_to_message != null) {
                data.price_to_message = this.price_to_message;
            }
            if (this.created != null) {
                data.created = this.created;
            }
            if (this.updated != null) {
                data.updated = this.updated;
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (this.id !== undefined)
                writer.writeInt64(1, this.id);
            if (typeof this.did === "string" && this.did.length)
                writer.writeString(2, this.did);
            if (typeof this.username === "string" && this.username.length)
                writer.writeString(3, this.username);
            if (typeof this.name === "string" && this.name.length)
                writer.writeString(4, this.name);
            if (typeof this.email === "string" && this.email.length)
                writer.writeString(5, this.email);
            if (typeof this.bio === "string" && this.bio.length)
                writer.writeString(6, this.bio);
            if (typeof this.img === "string" && this.img.length)
                writer.writeString(7, this.img);
            if (this.price_to_message !== undefined)
                writer.writeInt64(8, this.price_to_message);
            if (this.created !== undefined)
                writer.writeInt64(9, this.created);
            if (this.updated !== undefined)
                writer.writeInt64(10, this.updated);
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): User {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new User();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        message.id = reader.readInt64();
                        break;
                    case 2:
                        message.did = reader.readString();
                        break;
                    case 3:
                        message.username = reader.readString();
                        break;
                    case 4:
                        message.name = reader.readString();
                        break;
                    case 5:
                        message.email = reader.readString();
                        break;
                    case 6:
                        message.bio = reader.readString();
                        break;
                    case 7:
                        message.img = reader.readString();
                        break;
                    case 8:
                        message.price_to_message = reader.readInt64();
                        break;
                    case 9:
                        message.created = reader.readInt64();
                        break;
                    case 10:
                        message.updated = reader.readInt64();
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): User {
            return User.deserialize(bytes);
        }
    }
}
