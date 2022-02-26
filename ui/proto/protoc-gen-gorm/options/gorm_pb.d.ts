// package: gorm
// file: proto/protoc-gen-gorm/options/gorm.proto

import * as jspb from "google-protobuf";
import * as google_protobuf_descriptor_pb from "google-protobuf/google/protobuf/descriptor_pb";

export class GormFileOptions extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GormFileOptions.AsObject;
  static toObject(includeInstance: boolean, msg: GormFileOptions): GormFileOptions.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GormFileOptions, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GormFileOptions;
  static deserializeBinaryFromReader(message: GormFileOptions, reader: jspb.BinaryReader): GormFileOptions;
}

export namespace GormFileOptions {
  export type AsObject = {
  }
}

export class GormMessageOptions extends jspb.Message {
  hasOrmable(): boolean;
  clearOrmable(): void;
  getOrmable(): boolean | undefined;
  setOrmable(value: boolean): void;

  clearIncludeList(): void;
  getIncludeList(): Array<ExtraField>;
  setIncludeList(value: Array<ExtraField>): void;
  addInclude(value?: ExtraField, index?: number): ExtraField;

  hasTable(): boolean;
  clearTable(): void;
  getTable(): string | undefined;
  setTable(value: string): void;

  hasMultiAccount(): boolean;
  clearMultiAccount(): void;
  getMultiAccount(): boolean | undefined;
  setMultiAccount(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GormMessageOptions.AsObject;
  static toObject(includeInstance: boolean, msg: GormMessageOptions): GormMessageOptions.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GormMessageOptions, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GormMessageOptions;
  static deserializeBinaryFromReader(message: GormMessageOptions, reader: jspb.BinaryReader): GormMessageOptions;
}

export namespace GormMessageOptions {
  export type AsObject = {
    ormable?: boolean,
    includeList: Array<ExtraField.AsObject>,
    table?: string,
    multiAccount?: boolean,
  }
}

export class ExtraField extends jspb.Message {
  hasType(): boolean;
  clearType(): void;
  getType(): string | undefined;
  setType(value: string): void;

  hasName(): boolean;
  clearName(): void;
  getName(): string | undefined;
  setName(value: string): void;

  hasTag(): boolean;
  clearTag(): void;
  getTag(): GormTag | undefined;
  setTag(value?: GormTag): void;

  hasPackage(): boolean;
  clearPackage(): void;
  getPackage(): string | undefined;
  setPackage(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ExtraField.AsObject;
  static toObject(includeInstance: boolean, msg: ExtraField): ExtraField.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ExtraField, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ExtraField;
  static deserializeBinaryFromReader(message: ExtraField, reader: jspb.BinaryReader): ExtraField;
}

export namespace ExtraField {
  export type AsObject = {
    type?: string,
    name?: string,
    tag?: GormTag.AsObject,
    pb_package?: string,
  }
}

export class GormFieldOptions extends jspb.Message {
  hasTag(): boolean;
  clearTag(): void;
  getTag(): GormTag | undefined;
  setTag(value?: GormTag): void;

  hasDrop(): boolean;
  clearDrop(): void;
  getDrop(): boolean | undefined;
  setDrop(value: boolean): void;

  hasHasOne(): boolean;
  clearHasOne(): void;
  getHasOne(): HasOneOptions | undefined;
  setHasOne(value?: HasOneOptions): void;

  hasBelongsTo(): boolean;
  clearBelongsTo(): void;
  getBelongsTo(): BelongsToOptions | undefined;
  setBelongsTo(value?: BelongsToOptions): void;

  hasHasMany(): boolean;
  clearHasMany(): void;
  getHasMany(): HasManyOptions | undefined;
  setHasMany(value?: HasManyOptions): void;

  hasManyToMany(): boolean;
  clearManyToMany(): void;
  getManyToMany(): ManyToManyOptions | undefined;
  setManyToMany(value?: ManyToManyOptions): void;

  hasReferenceOf(): boolean;
  clearReferenceOf(): void;
  getReferenceOf(): string | undefined;
  setReferenceOf(value: string): void;

  getAssociationCase(): GormFieldOptions.AssociationCase;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GormFieldOptions.AsObject;
  static toObject(includeInstance: boolean, msg: GormFieldOptions): GormFieldOptions.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GormFieldOptions, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GormFieldOptions;
  static deserializeBinaryFromReader(message: GormFieldOptions, reader: jspb.BinaryReader): GormFieldOptions;
}

export namespace GormFieldOptions {
  export type AsObject = {
    tag?: GormTag.AsObject,
    drop?: boolean,
    hasOne?: HasOneOptions.AsObject,
    belongsTo?: BelongsToOptions.AsObject,
    hasMany?: HasManyOptions.AsObject,
    manyToMany?: ManyToManyOptions.AsObject,
    referenceOf?: string,
  }

  export enum AssociationCase {
    ASSOCIATION_NOT_SET = 0,
    HAS_ONE = 3,
    BELONGS_TO = 4,
    HAS_MANY = 5,
    MANY_TO_MANY = 6,
  }
}

export class GormTag extends jspb.Message {
  hasColumn(): boolean;
  clearColumn(): void;
  getColumn(): string | undefined;
  setColumn(value: string): void;

  hasType(): boolean;
  clearType(): void;
  getType(): string | undefined;
  setType(value: string): void;

  hasSize(): boolean;
  clearSize(): void;
  getSize(): number | undefined;
  setSize(value: number): void;

  hasPrecision(): boolean;
  clearPrecision(): void;
  getPrecision(): number | undefined;
  setPrecision(value: number): void;

  hasPrimaryKey(): boolean;
  clearPrimaryKey(): void;
  getPrimaryKey(): boolean | undefined;
  setPrimaryKey(value: boolean): void;

  hasUnique(): boolean;
  clearUnique(): void;
  getUnique(): boolean | undefined;
  setUnique(value: boolean): void;

  hasDefault(): boolean;
  clearDefault(): void;
  getDefault(): string | undefined;
  setDefault(value: string): void;

  hasNotNull(): boolean;
  clearNotNull(): void;
  getNotNull(): boolean | undefined;
  setNotNull(value: boolean): void;

  hasAutoIncrement(): boolean;
  clearAutoIncrement(): void;
  getAutoIncrement(): boolean | undefined;
  setAutoIncrement(value: boolean): void;

  hasIndex(): boolean;
  clearIndex(): void;
  getIndex(): string | undefined;
  setIndex(value: string): void;

  hasUniqueIndex(): boolean;
  clearUniqueIndex(): void;
  getUniqueIndex(): string | undefined;
  setUniqueIndex(value: string): void;

  hasEmbedded(): boolean;
  clearEmbedded(): void;
  getEmbedded(): boolean | undefined;
  setEmbedded(value: boolean): void;

  hasEmbeddedPrefix(): boolean;
  clearEmbeddedPrefix(): void;
  getEmbeddedPrefix(): string | undefined;
  setEmbeddedPrefix(value: string): void;

  hasIgnore(): boolean;
  clearIgnore(): void;
  getIgnore(): boolean | undefined;
  setIgnore(value: boolean): void;

  hasForeignkey(): boolean;
  clearForeignkey(): void;
  getForeignkey(): string | undefined;
  setForeignkey(value: string): void;

  hasAssociationForeignkey(): boolean;
  clearAssociationForeignkey(): void;
  getAssociationForeignkey(): string | undefined;
  setAssociationForeignkey(value: string): void;

  hasManyToMany(): boolean;
  clearManyToMany(): void;
  getManyToMany(): string | undefined;
  setManyToMany(value: string): void;

  hasJointableForeignkey(): boolean;
  clearJointableForeignkey(): void;
  getJointableForeignkey(): string | undefined;
  setJointableForeignkey(value: string): void;

  hasAssociationJointableForeignkey(): boolean;
  clearAssociationJointableForeignkey(): void;
  getAssociationJointableForeignkey(): string | undefined;
  setAssociationJointableForeignkey(value: string): void;

  hasAssociationAutoupdate(): boolean;
  clearAssociationAutoupdate(): void;
  getAssociationAutoupdate(): boolean | undefined;
  setAssociationAutoupdate(value: boolean): void;

  hasAssociationAutocreate(): boolean;
  clearAssociationAutocreate(): void;
  getAssociationAutocreate(): boolean | undefined;
  setAssociationAutocreate(value: boolean): void;

  hasAssociationSaveReference(): boolean;
  clearAssociationSaveReference(): void;
  getAssociationSaveReference(): boolean | undefined;
  setAssociationSaveReference(value: boolean): void;

  hasPreload(): boolean;
  clearPreload(): void;
  getPreload(): boolean | undefined;
  setPreload(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GormTag.AsObject;
  static toObject(includeInstance: boolean, msg: GormTag): GormTag.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GormTag, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GormTag;
  static deserializeBinaryFromReader(message: GormTag, reader: jspb.BinaryReader): GormTag;
}

export namespace GormTag {
  export type AsObject = {
    column?: string,
    type?: string,
    size?: number,
    precision?: number,
    primaryKey?: boolean,
    unique?: boolean,
    pb_default?: string,
    notNull?: boolean,
    autoIncrement?: boolean,
    index?: string,
    uniqueIndex?: string,
    embedded?: boolean,
    embeddedPrefix?: string,
    ignore?: boolean,
    foreignkey?: string,
    associationForeignkey?: string,
    manyToMany?: string,
    jointableForeignkey?: string,
    associationJointableForeignkey?: string,
    associationAutoupdate?: boolean,
    associationAutocreate?: boolean,
    associationSaveReference?: boolean,
    preload?: boolean,
  }
}

export class HasOneOptions extends jspb.Message {
  hasForeignkey(): boolean;
  clearForeignkey(): void;
  getForeignkey(): string | undefined;
  setForeignkey(value: string): void;

  hasForeignkeyTag(): boolean;
  clearForeignkeyTag(): void;
  getForeignkeyTag(): GormTag | undefined;
  setForeignkeyTag(value?: GormTag): void;

  hasAssociationForeignkey(): boolean;
  clearAssociationForeignkey(): void;
  getAssociationForeignkey(): string | undefined;
  setAssociationForeignkey(value: string): void;

  hasAssociationAutoupdate(): boolean;
  clearAssociationAutoupdate(): void;
  getAssociationAutoupdate(): boolean | undefined;
  setAssociationAutoupdate(value: boolean): void;

  hasAssociationAutocreate(): boolean;
  clearAssociationAutocreate(): void;
  getAssociationAutocreate(): boolean | undefined;
  setAssociationAutocreate(value: boolean): void;

  hasAssociationSaveReference(): boolean;
  clearAssociationSaveReference(): void;
  getAssociationSaveReference(): boolean | undefined;
  setAssociationSaveReference(value: boolean): void;

  hasPreload(): boolean;
  clearPreload(): void;
  getPreload(): boolean | undefined;
  setPreload(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HasOneOptions.AsObject;
  static toObject(includeInstance: boolean, msg: HasOneOptions): HasOneOptions.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: HasOneOptions, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HasOneOptions;
  static deserializeBinaryFromReader(message: HasOneOptions, reader: jspb.BinaryReader): HasOneOptions;
}

export namespace HasOneOptions {
  export type AsObject = {
    foreignkey?: string,
    foreignkeyTag?: GormTag.AsObject,
    associationForeignkey?: string,
    associationAutoupdate?: boolean,
    associationAutocreate?: boolean,
    associationSaveReference?: boolean,
    preload?: boolean,
  }
}

export class BelongsToOptions extends jspb.Message {
  hasForeignkey(): boolean;
  clearForeignkey(): void;
  getForeignkey(): string | undefined;
  setForeignkey(value: string): void;

  hasForeignkeyTag(): boolean;
  clearForeignkeyTag(): void;
  getForeignkeyTag(): GormTag | undefined;
  setForeignkeyTag(value?: GormTag): void;

  hasAssociationForeignkey(): boolean;
  clearAssociationForeignkey(): void;
  getAssociationForeignkey(): string | undefined;
  setAssociationForeignkey(value: string): void;

  hasAssociationAutoupdate(): boolean;
  clearAssociationAutoupdate(): void;
  getAssociationAutoupdate(): boolean | undefined;
  setAssociationAutoupdate(value: boolean): void;

  hasAssociationAutocreate(): boolean;
  clearAssociationAutocreate(): void;
  getAssociationAutocreate(): boolean | undefined;
  setAssociationAutocreate(value: boolean): void;

  hasAssociationSaveReference(): boolean;
  clearAssociationSaveReference(): void;
  getAssociationSaveReference(): boolean | undefined;
  setAssociationSaveReference(value: boolean): void;

  hasPreload(): boolean;
  clearPreload(): void;
  getPreload(): boolean | undefined;
  setPreload(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): BelongsToOptions.AsObject;
  static toObject(includeInstance: boolean, msg: BelongsToOptions): BelongsToOptions.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: BelongsToOptions, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): BelongsToOptions;
  static deserializeBinaryFromReader(message: BelongsToOptions, reader: jspb.BinaryReader): BelongsToOptions;
}

export namespace BelongsToOptions {
  export type AsObject = {
    foreignkey?: string,
    foreignkeyTag?: GormTag.AsObject,
    associationForeignkey?: string,
    associationAutoupdate?: boolean,
    associationAutocreate?: boolean,
    associationSaveReference?: boolean,
    preload?: boolean,
  }
}

export class HasManyOptions extends jspb.Message {
  hasForeignkey(): boolean;
  clearForeignkey(): void;
  getForeignkey(): string | undefined;
  setForeignkey(value: string): void;

  hasForeignkeyTag(): boolean;
  clearForeignkeyTag(): void;
  getForeignkeyTag(): GormTag | undefined;
  setForeignkeyTag(value?: GormTag): void;

  hasAssociationForeignkey(): boolean;
  clearAssociationForeignkey(): void;
  getAssociationForeignkey(): string | undefined;
  setAssociationForeignkey(value: string): void;

  hasPositionField(): boolean;
  clearPositionField(): void;
  getPositionField(): string | undefined;
  setPositionField(value: string): void;

  hasPositionFieldTag(): boolean;
  clearPositionFieldTag(): void;
  getPositionFieldTag(): GormTag | undefined;
  setPositionFieldTag(value?: GormTag): void;

  hasAssociationAutoupdate(): boolean;
  clearAssociationAutoupdate(): void;
  getAssociationAutoupdate(): boolean | undefined;
  setAssociationAutoupdate(value: boolean): void;

  hasAssociationAutocreate(): boolean;
  clearAssociationAutocreate(): void;
  getAssociationAutocreate(): boolean | undefined;
  setAssociationAutocreate(value: boolean): void;

  hasAssociationSaveReference(): boolean;
  clearAssociationSaveReference(): void;
  getAssociationSaveReference(): boolean | undefined;
  setAssociationSaveReference(value: boolean): void;

  hasPreload(): boolean;
  clearPreload(): void;
  getPreload(): boolean | undefined;
  setPreload(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): HasManyOptions.AsObject;
  static toObject(includeInstance: boolean, msg: HasManyOptions): HasManyOptions.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: HasManyOptions, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): HasManyOptions;
  static deserializeBinaryFromReader(message: HasManyOptions, reader: jspb.BinaryReader): HasManyOptions;
}

export namespace HasManyOptions {
  export type AsObject = {
    foreignkey?: string,
    foreignkeyTag?: GormTag.AsObject,
    associationForeignkey?: string,
    positionField?: string,
    positionFieldTag?: GormTag.AsObject,
    associationAutoupdate?: boolean,
    associationAutocreate?: boolean,
    associationSaveReference?: boolean,
    preload?: boolean,
  }
}

export class ManyToManyOptions extends jspb.Message {
  hasJointable(): boolean;
  clearJointable(): void;
  getJointable(): string | undefined;
  setJointable(value: string): void;

  hasForeignkey(): boolean;
  clearForeignkey(): void;
  getForeignkey(): string | undefined;
  setForeignkey(value: string): void;

  hasJointableForeignkey(): boolean;
  clearJointableForeignkey(): void;
  getJointableForeignkey(): string | undefined;
  setJointableForeignkey(value: string): void;

  hasAssociationForeignkey(): boolean;
  clearAssociationForeignkey(): void;
  getAssociationForeignkey(): string | undefined;
  setAssociationForeignkey(value: string): void;

  hasAssociationJointableForeignkey(): boolean;
  clearAssociationJointableForeignkey(): void;
  getAssociationJointableForeignkey(): string | undefined;
  setAssociationJointableForeignkey(value: string): void;

  hasAssociationAutoupdate(): boolean;
  clearAssociationAutoupdate(): void;
  getAssociationAutoupdate(): boolean | undefined;
  setAssociationAutoupdate(value: boolean): void;

  hasAssociationAutocreate(): boolean;
  clearAssociationAutocreate(): void;
  getAssociationAutocreate(): boolean | undefined;
  setAssociationAutocreate(value: boolean): void;

  hasAssociationSaveReference(): boolean;
  clearAssociationSaveReference(): void;
  getAssociationSaveReference(): boolean | undefined;
  setAssociationSaveReference(value: boolean): void;

  hasPreload(): boolean;
  clearPreload(): void;
  getPreload(): boolean | undefined;
  setPreload(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ManyToManyOptions.AsObject;
  static toObject(includeInstance: boolean, msg: ManyToManyOptions): ManyToManyOptions.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ManyToManyOptions, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ManyToManyOptions;
  static deserializeBinaryFromReader(message: ManyToManyOptions, reader: jspb.BinaryReader): ManyToManyOptions;
}

export namespace ManyToManyOptions {
  export type AsObject = {
    jointable?: string,
    foreignkey?: string,
    jointableForeignkey?: string,
    associationForeignkey?: string,
    associationJointableForeignkey?: string,
    associationAutoupdate?: boolean,
    associationAutocreate?: boolean,
    associationSaveReference?: boolean,
    preload?: boolean,
  }
}

export class AutoServerOptions extends jspb.Message {
  hasAutogen(): boolean;
  clearAutogen(): void;
  getAutogen(): boolean | undefined;
  setAutogen(value: boolean): void;

  hasTxnMiddleware(): boolean;
  clearTxnMiddleware(): void;
  getTxnMiddleware(): boolean | undefined;
  setTxnMiddleware(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AutoServerOptions.AsObject;
  static toObject(includeInstance: boolean, msg: AutoServerOptions): AutoServerOptions.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: AutoServerOptions, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AutoServerOptions;
  static deserializeBinaryFromReader(message: AutoServerOptions, reader: jspb.BinaryReader): AutoServerOptions;
}

export namespace AutoServerOptions {
  export type AsObject = {
    autogen?: boolean,
    txnMiddleware?: boolean,
  }
}

export class MethodOptions extends jspb.Message {
  hasObjectType(): boolean;
  clearObjectType(): void;
  getObjectType(): string | undefined;
  setObjectType(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): MethodOptions.AsObject;
  static toObject(includeInstance: boolean, msg: MethodOptions): MethodOptions.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: MethodOptions, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): MethodOptions;
  static deserializeBinaryFromReader(message: MethodOptions, reader: jspb.BinaryReader): MethodOptions;
}

export namespace MethodOptions {
  export type AsObject = {
    objectType?: string,
  }
}

  export const fileOpts: jspb.ExtensionFieldInfo<GormFileOptions>;

  export const opts: jspb.ExtensionFieldInfo<GormMessageOptions>;

  export const field: jspb.ExtensionFieldInfo<GormFieldOptions>;

  export const server: jspb.ExtensionFieldInfo<AutoServerOptions>;

  export const method: jspb.ExtensionFieldInfo<MethodOptions>;

