// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mail/mail_transaction.proto

package mail;

public final class MailTransactionOuterClass {
  private MailTransactionOuterClass() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  public interface MailTransactionOrBuilder extends
      // @@protoc_insertion_point(interface_extends:mail.MailTransaction)
      com.google.protobuf.MessageOrBuilder {

    /**
     * <code>string id = 1;</code>
     * @return The id.
     */
    java.lang.String getId();
    /**
     * <code>string id = 1;</code>
     * @return The bytes for id.
     */
    com.google.protobuf.ByteString
        getIdBytes();

    /**
     * <code>int64 tx_created_at = 3;</code>
     * @return The txCreatedAt.
     */
    long getTxCreatedAt();

    /**
     * <code>repeated .mail.Indicator indicators = 4;</code>
     */
    java.util.List<mail.IndicatorOuterClass.Indicator> 
        getIndicatorsList();
    /**
     * <code>repeated .mail.Indicator indicators = 4;</code>
     */
    mail.IndicatorOuterClass.Indicator getIndicators(int index);
    /**
     * <code>repeated .mail.Indicator indicators = 4;</code>
     */
    int getIndicatorsCount();
    /**
     * <code>repeated .mail.Indicator indicators = 4;</code>
     */
    java.util.List<? extends mail.IndicatorOuterClass.IndicatorOrBuilder> 
        getIndicatorsOrBuilderList();
    /**
     * <code>repeated .mail.Indicator indicators = 4;</code>
     */
    mail.IndicatorOuterClass.IndicatorOrBuilder getIndicatorsOrBuilder(
        int index);
  }
  /**
   * Protobuf type {@code mail.MailTransaction}
   */
  public  static final class MailTransaction extends
      com.google.protobuf.GeneratedMessageV3 implements
      // @@protoc_insertion_point(message_implements:mail.MailTransaction)
      MailTransactionOrBuilder {
  private static final long serialVersionUID = 0L;
    // Use MailTransaction.newBuilder() to construct.
    private MailTransaction(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
      super(builder);
    }
    private MailTransaction() {
      id_ = "";
      indicators_ = java.util.Collections.emptyList();
    }

    @java.lang.Override
    @SuppressWarnings({"unused"})
    protected java.lang.Object newInstance(
        UnusedPrivateParameter unused) {
      return new MailTransaction();
    }

    @java.lang.Override
    public final com.google.protobuf.UnknownFieldSet
    getUnknownFields() {
      return this.unknownFields;
    }
    private MailTransaction(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      this();
      if (extensionRegistry == null) {
        throw new java.lang.NullPointerException();
      }
      int mutable_bitField0_ = 0;
      com.google.protobuf.UnknownFieldSet.Builder unknownFields =
          com.google.protobuf.UnknownFieldSet.newBuilder();
      try {
        boolean done = false;
        while (!done) {
          int tag = input.readTag();
          switch (tag) {
            case 0:
              done = true;
              break;
            case 10: {
              java.lang.String s = input.readStringRequireUtf8();

              id_ = s;
              break;
            }
            case 24: {

              txCreatedAt_ = input.readInt64();
              break;
            }
            case 34: {
              if (!((mutable_bitField0_ & 0x00000001) != 0)) {
                indicators_ = new java.util.ArrayList<mail.IndicatorOuterClass.Indicator>();
                mutable_bitField0_ |= 0x00000001;
              }
              indicators_.add(
                  input.readMessage(mail.IndicatorOuterClass.Indicator.parser(), extensionRegistry));
              break;
            }
            default: {
              if (!parseUnknownField(
                  input, unknownFields, extensionRegistry, tag)) {
                done = true;
              }
              break;
            }
          }
        }
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        throw e.setUnfinishedMessage(this);
      } catch (java.io.IOException e) {
        throw new com.google.protobuf.InvalidProtocolBufferException(
            e).setUnfinishedMessage(this);
      } finally {
        if (((mutable_bitField0_ & 0x00000001) != 0)) {
          indicators_ = java.util.Collections.unmodifiableList(indicators_);
        }
        this.unknownFields = unknownFields.build();
        makeExtensionsImmutable();
      }
    }
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return mail.MailTransactionOuterClass.internal_static_mail_MailTransaction_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return mail.MailTransactionOuterClass.internal_static_mail_MailTransaction_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              mail.MailTransactionOuterClass.MailTransaction.class, mail.MailTransactionOuterClass.MailTransaction.Builder.class);
    }

    public static final int ID_FIELD_NUMBER = 1;
    private volatile java.lang.Object id_;
    /**
     * <code>string id = 1;</code>
     * @return The id.
     */
    public java.lang.String getId() {
      java.lang.Object ref = id_;
      if (ref instanceof java.lang.String) {
        return (java.lang.String) ref;
      } else {
        com.google.protobuf.ByteString bs = 
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        id_ = s;
        return s;
      }
    }
    /**
     * <code>string id = 1;</code>
     * @return The bytes for id.
     */
    public com.google.protobuf.ByteString
        getIdBytes() {
      java.lang.Object ref = id_;
      if (ref instanceof java.lang.String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        id_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }

    public static final int TX_CREATED_AT_FIELD_NUMBER = 3;
    private long txCreatedAt_;
    /**
     * <code>int64 tx_created_at = 3;</code>
     * @return The txCreatedAt.
     */
    public long getTxCreatedAt() {
      return txCreatedAt_;
    }

    public static final int INDICATORS_FIELD_NUMBER = 4;
    private java.util.List<mail.IndicatorOuterClass.Indicator> indicators_;
    /**
     * <code>repeated .mail.Indicator indicators = 4;</code>
     */
    public java.util.List<mail.IndicatorOuterClass.Indicator> getIndicatorsList() {
      return indicators_;
    }
    /**
     * <code>repeated .mail.Indicator indicators = 4;</code>
     */
    public java.util.List<? extends mail.IndicatorOuterClass.IndicatorOrBuilder> 
        getIndicatorsOrBuilderList() {
      return indicators_;
    }
    /**
     * <code>repeated .mail.Indicator indicators = 4;</code>
     */
    public int getIndicatorsCount() {
      return indicators_.size();
    }
    /**
     * <code>repeated .mail.Indicator indicators = 4;</code>
     */
    public mail.IndicatorOuterClass.Indicator getIndicators(int index) {
      return indicators_.get(index);
    }
    /**
     * <code>repeated .mail.Indicator indicators = 4;</code>
     */
    public mail.IndicatorOuterClass.IndicatorOrBuilder getIndicatorsOrBuilder(
        int index) {
      return indicators_.get(index);
    }

    private byte memoizedIsInitialized = -1;
    @java.lang.Override
    public final boolean isInitialized() {
      byte isInitialized = memoizedIsInitialized;
      if (isInitialized == 1) return true;
      if (isInitialized == 0) return false;

      memoizedIsInitialized = 1;
      return true;
    }

    @java.lang.Override
    public void writeTo(com.google.protobuf.CodedOutputStream output)
                        throws java.io.IOException {
      if (!getIdBytes().isEmpty()) {
        com.google.protobuf.GeneratedMessageV3.writeString(output, 1, id_);
      }
      if (txCreatedAt_ != 0L) {
        output.writeInt64(3, txCreatedAt_);
      }
      for (int i = 0; i < indicators_.size(); i++) {
        output.writeMessage(4, indicators_.get(i));
      }
      unknownFields.writeTo(output);
    }

    @java.lang.Override
    public int getSerializedSize() {
      int size = memoizedSize;
      if (size != -1) return size;

      size = 0;
      if (!getIdBytes().isEmpty()) {
        size += com.google.protobuf.GeneratedMessageV3.computeStringSize(1, id_);
      }
      if (txCreatedAt_ != 0L) {
        size += com.google.protobuf.CodedOutputStream
          .computeInt64Size(3, txCreatedAt_);
      }
      for (int i = 0; i < indicators_.size(); i++) {
        size += com.google.protobuf.CodedOutputStream
          .computeMessageSize(4, indicators_.get(i));
      }
      size += unknownFields.getSerializedSize();
      memoizedSize = size;
      return size;
    }

    @java.lang.Override
    public boolean equals(final java.lang.Object obj) {
      if (obj == this) {
       return true;
      }
      if (!(obj instanceof mail.MailTransactionOuterClass.MailTransaction)) {
        return super.equals(obj);
      }
      mail.MailTransactionOuterClass.MailTransaction other = (mail.MailTransactionOuterClass.MailTransaction) obj;

      if (!getId()
          .equals(other.getId())) return false;
      if (getTxCreatedAt()
          != other.getTxCreatedAt()) return false;
      if (!getIndicatorsList()
          .equals(other.getIndicatorsList())) return false;
      if (!unknownFields.equals(other.unknownFields)) return false;
      return true;
    }

    @java.lang.Override
    public int hashCode() {
      if (memoizedHashCode != 0) {
        return memoizedHashCode;
      }
      int hash = 41;
      hash = (19 * hash) + getDescriptor().hashCode();
      hash = (37 * hash) + ID_FIELD_NUMBER;
      hash = (53 * hash) + getId().hashCode();
      hash = (37 * hash) + TX_CREATED_AT_FIELD_NUMBER;
      hash = (53 * hash) + com.google.protobuf.Internal.hashLong(
          getTxCreatedAt());
      if (getIndicatorsCount() > 0) {
        hash = (37 * hash) + INDICATORS_FIELD_NUMBER;
        hash = (53 * hash) + getIndicatorsList().hashCode();
      }
      hash = (29 * hash) + unknownFields.hashCode();
      memoizedHashCode = hash;
      return hash;
    }

    public static mail.MailTransactionOuterClass.MailTransaction parseFrom(
        java.nio.ByteBuffer data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static mail.MailTransactionOuterClass.MailTransaction parseFrom(
        java.nio.ByteBuffer data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static mail.MailTransactionOuterClass.MailTransaction parseFrom(
        com.google.protobuf.ByteString data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static mail.MailTransactionOuterClass.MailTransaction parseFrom(
        com.google.protobuf.ByteString data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static mail.MailTransactionOuterClass.MailTransaction parseFrom(byte[] data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static mail.MailTransactionOuterClass.MailTransaction parseFrom(
        byte[] data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static mail.MailTransactionOuterClass.MailTransaction parseFrom(java.io.InputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input);
    }
    public static mail.MailTransactionOuterClass.MailTransaction parseFrom(
        java.io.InputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input, extensionRegistry);
    }
    public static mail.MailTransactionOuterClass.MailTransaction parseDelimitedFrom(java.io.InputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseDelimitedWithIOException(PARSER, input);
    }
    public static mail.MailTransactionOuterClass.MailTransaction parseDelimitedFrom(
        java.io.InputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
    }
    public static mail.MailTransactionOuterClass.MailTransaction parseFrom(
        com.google.protobuf.CodedInputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input);
    }
    public static mail.MailTransactionOuterClass.MailTransaction parseFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input, extensionRegistry);
    }

    @java.lang.Override
    public Builder newBuilderForType() { return newBuilder(); }
    public static Builder newBuilder() {
      return DEFAULT_INSTANCE.toBuilder();
    }
    public static Builder newBuilder(mail.MailTransactionOuterClass.MailTransaction prototype) {
      return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
    }
    @java.lang.Override
    public Builder toBuilder() {
      return this == DEFAULT_INSTANCE
          ? new Builder() : new Builder().mergeFrom(this);
    }

    @java.lang.Override
    protected Builder newBuilderForType(
        com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      Builder builder = new Builder(parent);
      return builder;
    }
    /**
     * Protobuf type {@code mail.MailTransaction}
     */
    public static final class Builder extends
        com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
        // @@protoc_insertion_point(builder_implements:mail.MailTransaction)
        mail.MailTransactionOuterClass.MailTransactionOrBuilder {
      public static final com.google.protobuf.Descriptors.Descriptor
          getDescriptor() {
        return mail.MailTransactionOuterClass.internal_static_mail_MailTransaction_descriptor;
      }

      @java.lang.Override
      protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
          internalGetFieldAccessorTable() {
        return mail.MailTransactionOuterClass.internal_static_mail_MailTransaction_fieldAccessorTable
            .ensureFieldAccessorsInitialized(
                mail.MailTransactionOuterClass.MailTransaction.class, mail.MailTransactionOuterClass.MailTransaction.Builder.class);
      }

      // Construct using mail.MailTransactionOuterClass.MailTransaction.newBuilder()
      private Builder() {
        maybeForceBuilderInitialization();
      }

      private Builder(
          com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
        super(parent);
        maybeForceBuilderInitialization();
      }
      private void maybeForceBuilderInitialization() {
        if (com.google.protobuf.GeneratedMessageV3
                .alwaysUseFieldBuilders) {
          getIndicatorsFieldBuilder();
        }
      }
      @java.lang.Override
      public Builder clear() {
        super.clear();
        id_ = "";

        txCreatedAt_ = 0L;

        if (indicatorsBuilder_ == null) {
          indicators_ = java.util.Collections.emptyList();
          bitField0_ = (bitField0_ & ~0x00000001);
        } else {
          indicatorsBuilder_.clear();
        }
        return this;
      }

      @java.lang.Override
      public com.google.protobuf.Descriptors.Descriptor
          getDescriptorForType() {
        return mail.MailTransactionOuterClass.internal_static_mail_MailTransaction_descriptor;
      }

      @java.lang.Override
      public mail.MailTransactionOuterClass.MailTransaction getDefaultInstanceForType() {
        return mail.MailTransactionOuterClass.MailTransaction.getDefaultInstance();
      }

      @java.lang.Override
      public mail.MailTransactionOuterClass.MailTransaction build() {
        mail.MailTransactionOuterClass.MailTransaction result = buildPartial();
        if (!result.isInitialized()) {
          throw newUninitializedMessageException(result);
        }
        return result;
      }

      @java.lang.Override
      public mail.MailTransactionOuterClass.MailTransaction buildPartial() {
        mail.MailTransactionOuterClass.MailTransaction result = new mail.MailTransactionOuterClass.MailTransaction(this);
        int from_bitField0_ = bitField0_;
        result.id_ = id_;
        result.txCreatedAt_ = txCreatedAt_;
        if (indicatorsBuilder_ == null) {
          if (((bitField0_ & 0x00000001) != 0)) {
            indicators_ = java.util.Collections.unmodifiableList(indicators_);
            bitField0_ = (bitField0_ & ~0x00000001);
          }
          result.indicators_ = indicators_;
        } else {
          result.indicators_ = indicatorsBuilder_.build();
        }
        onBuilt();
        return result;
      }

      @java.lang.Override
      public Builder clone() {
        return super.clone();
      }
      @java.lang.Override
      public Builder setField(
          com.google.protobuf.Descriptors.FieldDescriptor field,
          java.lang.Object value) {
        return super.setField(field, value);
      }
      @java.lang.Override
      public Builder clearField(
          com.google.protobuf.Descriptors.FieldDescriptor field) {
        return super.clearField(field);
      }
      @java.lang.Override
      public Builder clearOneof(
          com.google.protobuf.Descriptors.OneofDescriptor oneof) {
        return super.clearOneof(oneof);
      }
      @java.lang.Override
      public Builder setRepeatedField(
          com.google.protobuf.Descriptors.FieldDescriptor field,
          int index, java.lang.Object value) {
        return super.setRepeatedField(field, index, value);
      }
      @java.lang.Override
      public Builder addRepeatedField(
          com.google.protobuf.Descriptors.FieldDescriptor field,
          java.lang.Object value) {
        return super.addRepeatedField(field, value);
      }
      @java.lang.Override
      public Builder mergeFrom(com.google.protobuf.Message other) {
        if (other instanceof mail.MailTransactionOuterClass.MailTransaction) {
          return mergeFrom((mail.MailTransactionOuterClass.MailTransaction)other);
        } else {
          super.mergeFrom(other);
          return this;
        }
      }

      public Builder mergeFrom(mail.MailTransactionOuterClass.MailTransaction other) {
        if (other == mail.MailTransactionOuterClass.MailTransaction.getDefaultInstance()) return this;
        if (!other.getId().isEmpty()) {
          id_ = other.id_;
          onChanged();
        }
        if (other.getTxCreatedAt() != 0L) {
          setTxCreatedAt(other.getTxCreatedAt());
        }
        if (indicatorsBuilder_ == null) {
          if (!other.indicators_.isEmpty()) {
            if (indicators_.isEmpty()) {
              indicators_ = other.indicators_;
              bitField0_ = (bitField0_ & ~0x00000001);
            } else {
              ensureIndicatorsIsMutable();
              indicators_.addAll(other.indicators_);
            }
            onChanged();
          }
        } else {
          if (!other.indicators_.isEmpty()) {
            if (indicatorsBuilder_.isEmpty()) {
              indicatorsBuilder_.dispose();
              indicatorsBuilder_ = null;
              indicators_ = other.indicators_;
              bitField0_ = (bitField0_ & ~0x00000001);
              indicatorsBuilder_ = 
                com.google.protobuf.GeneratedMessageV3.alwaysUseFieldBuilders ?
                   getIndicatorsFieldBuilder() : null;
            } else {
              indicatorsBuilder_.addAllMessages(other.indicators_);
            }
          }
        }
        this.mergeUnknownFields(other.unknownFields);
        onChanged();
        return this;
      }

      @java.lang.Override
      public final boolean isInitialized() {
        return true;
      }

      @java.lang.Override
      public Builder mergeFrom(
          com.google.protobuf.CodedInputStream input,
          com.google.protobuf.ExtensionRegistryLite extensionRegistry)
          throws java.io.IOException {
        mail.MailTransactionOuterClass.MailTransaction parsedMessage = null;
        try {
          parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
        } catch (com.google.protobuf.InvalidProtocolBufferException e) {
          parsedMessage = (mail.MailTransactionOuterClass.MailTransaction) e.getUnfinishedMessage();
          throw e.unwrapIOException();
        } finally {
          if (parsedMessage != null) {
            mergeFrom(parsedMessage);
          }
        }
        return this;
      }
      private int bitField0_;

      private java.lang.Object id_ = "";
      /**
       * <code>string id = 1;</code>
       * @return The id.
       */
      public java.lang.String getId() {
        java.lang.Object ref = id_;
        if (!(ref instanceof java.lang.String)) {
          com.google.protobuf.ByteString bs =
              (com.google.protobuf.ByteString) ref;
          java.lang.String s = bs.toStringUtf8();
          id_ = s;
          return s;
        } else {
          return (java.lang.String) ref;
        }
      }
      /**
       * <code>string id = 1;</code>
       * @return The bytes for id.
       */
      public com.google.protobuf.ByteString
          getIdBytes() {
        java.lang.Object ref = id_;
        if (ref instanceof String) {
          com.google.protobuf.ByteString b = 
              com.google.protobuf.ByteString.copyFromUtf8(
                  (java.lang.String) ref);
          id_ = b;
          return b;
        } else {
          return (com.google.protobuf.ByteString) ref;
        }
      }
      /**
       * <code>string id = 1;</code>
       * @param value The id to set.
       * @return This builder for chaining.
       */
      public Builder setId(
          java.lang.String value) {
        if (value == null) {
    throw new NullPointerException();
  }
  
        id_ = value;
        onChanged();
        return this;
      }
      /**
       * <code>string id = 1;</code>
       * @return This builder for chaining.
       */
      public Builder clearId() {
        
        id_ = getDefaultInstance().getId();
        onChanged();
        return this;
      }
      /**
       * <code>string id = 1;</code>
       * @param value The bytes for id to set.
       * @return This builder for chaining.
       */
      public Builder setIdBytes(
          com.google.protobuf.ByteString value) {
        if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
        
        id_ = value;
        onChanged();
        return this;
      }

      private long txCreatedAt_ ;
      /**
       * <code>int64 tx_created_at = 3;</code>
       * @return The txCreatedAt.
       */
      public long getTxCreatedAt() {
        return txCreatedAt_;
      }
      /**
       * <code>int64 tx_created_at = 3;</code>
       * @param value The txCreatedAt to set.
       * @return This builder for chaining.
       */
      public Builder setTxCreatedAt(long value) {
        
        txCreatedAt_ = value;
        onChanged();
        return this;
      }
      /**
       * <code>int64 tx_created_at = 3;</code>
       * @return This builder for chaining.
       */
      public Builder clearTxCreatedAt() {
        
        txCreatedAt_ = 0L;
        onChanged();
        return this;
      }

      private java.util.List<mail.IndicatorOuterClass.Indicator> indicators_ =
        java.util.Collections.emptyList();
      private void ensureIndicatorsIsMutable() {
        if (!((bitField0_ & 0x00000001) != 0)) {
          indicators_ = new java.util.ArrayList<mail.IndicatorOuterClass.Indicator>(indicators_);
          bitField0_ |= 0x00000001;
         }
      }

      private com.google.protobuf.RepeatedFieldBuilderV3<
          mail.IndicatorOuterClass.Indicator, mail.IndicatorOuterClass.Indicator.Builder, mail.IndicatorOuterClass.IndicatorOrBuilder> indicatorsBuilder_;

      /**
       * <code>repeated .mail.Indicator indicators = 4;</code>
       */
      public java.util.List<mail.IndicatorOuterClass.Indicator> getIndicatorsList() {
        if (indicatorsBuilder_ == null) {
          return java.util.Collections.unmodifiableList(indicators_);
        } else {
          return indicatorsBuilder_.getMessageList();
        }
      }
      /**
       * <code>repeated .mail.Indicator indicators = 4;</code>
       */
      public int getIndicatorsCount() {
        if (indicatorsBuilder_ == null) {
          return indicators_.size();
        } else {
          return indicatorsBuilder_.getCount();
        }
      }
      /**
       * <code>repeated .mail.Indicator indicators = 4;</code>
       */
      public mail.IndicatorOuterClass.Indicator getIndicators(int index) {
        if (indicatorsBuilder_ == null) {
          return indicators_.get(index);
        } else {
          return indicatorsBuilder_.getMessage(index);
        }
      }
      /**
       * <code>repeated .mail.Indicator indicators = 4;</code>
       */
      public Builder setIndicators(
          int index, mail.IndicatorOuterClass.Indicator value) {
        if (indicatorsBuilder_ == null) {
          if (value == null) {
            throw new NullPointerException();
          }
          ensureIndicatorsIsMutable();
          indicators_.set(index, value);
          onChanged();
        } else {
          indicatorsBuilder_.setMessage(index, value);
        }
        return this;
      }
      /**
       * <code>repeated .mail.Indicator indicators = 4;</code>
       */
      public Builder setIndicators(
          int index, mail.IndicatorOuterClass.Indicator.Builder builderForValue) {
        if (indicatorsBuilder_ == null) {
          ensureIndicatorsIsMutable();
          indicators_.set(index, builderForValue.build());
          onChanged();
        } else {
          indicatorsBuilder_.setMessage(index, builderForValue.build());
        }
        return this;
      }
      /**
       * <code>repeated .mail.Indicator indicators = 4;</code>
       */
      public Builder addIndicators(mail.IndicatorOuterClass.Indicator value) {
        if (indicatorsBuilder_ == null) {
          if (value == null) {
            throw new NullPointerException();
          }
          ensureIndicatorsIsMutable();
          indicators_.add(value);
          onChanged();
        } else {
          indicatorsBuilder_.addMessage(value);
        }
        return this;
      }
      /**
       * <code>repeated .mail.Indicator indicators = 4;</code>
       */
      public Builder addIndicators(
          int index, mail.IndicatorOuterClass.Indicator value) {
        if (indicatorsBuilder_ == null) {
          if (value == null) {
            throw new NullPointerException();
          }
          ensureIndicatorsIsMutable();
          indicators_.add(index, value);
          onChanged();
        } else {
          indicatorsBuilder_.addMessage(index, value);
        }
        return this;
      }
      /**
       * <code>repeated .mail.Indicator indicators = 4;</code>
       */
      public Builder addIndicators(
          mail.IndicatorOuterClass.Indicator.Builder builderForValue) {
        if (indicatorsBuilder_ == null) {
          ensureIndicatorsIsMutable();
          indicators_.add(builderForValue.build());
          onChanged();
        } else {
          indicatorsBuilder_.addMessage(builderForValue.build());
        }
        return this;
      }
      /**
       * <code>repeated .mail.Indicator indicators = 4;</code>
       */
      public Builder addIndicators(
          int index, mail.IndicatorOuterClass.Indicator.Builder builderForValue) {
        if (indicatorsBuilder_ == null) {
          ensureIndicatorsIsMutable();
          indicators_.add(index, builderForValue.build());
          onChanged();
        } else {
          indicatorsBuilder_.addMessage(index, builderForValue.build());
        }
        return this;
      }
      /**
       * <code>repeated .mail.Indicator indicators = 4;</code>
       */
      public Builder addAllIndicators(
          java.lang.Iterable<? extends mail.IndicatorOuterClass.Indicator> values) {
        if (indicatorsBuilder_ == null) {
          ensureIndicatorsIsMutable();
          com.google.protobuf.AbstractMessageLite.Builder.addAll(
              values, indicators_);
          onChanged();
        } else {
          indicatorsBuilder_.addAllMessages(values);
        }
        return this;
      }
      /**
       * <code>repeated .mail.Indicator indicators = 4;</code>
       */
      public Builder clearIndicators() {
        if (indicatorsBuilder_ == null) {
          indicators_ = java.util.Collections.emptyList();
          bitField0_ = (bitField0_ & ~0x00000001);
          onChanged();
        } else {
          indicatorsBuilder_.clear();
        }
        return this;
      }
      /**
       * <code>repeated .mail.Indicator indicators = 4;</code>
       */
      public Builder removeIndicators(int index) {
        if (indicatorsBuilder_ == null) {
          ensureIndicatorsIsMutable();
          indicators_.remove(index);
          onChanged();
        } else {
          indicatorsBuilder_.remove(index);
        }
        return this;
      }
      /**
       * <code>repeated .mail.Indicator indicators = 4;</code>
       */
      public mail.IndicatorOuterClass.Indicator.Builder getIndicatorsBuilder(
          int index) {
        return getIndicatorsFieldBuilder().getBuilder(index);
      }
      /**
       * <code>repeated .mail.Indicator indicators = 4;</code>
       */
      public mail.IndicatorOuterClass.IndicatorOrBuilder getIndicatorsOrBuilder(
          int index) {
        if (indicatorsBuilder_ == null) {
          return indicators_.get(index);  } else {
          return indicatorsBuilder_.getMessageOrBuilder(index);
        }
      }
      /**
       * <code>repeated .mail.Indicator indicators = 4;</code>
       */
      public java.util.List<? extends mail.IndicatorOuterClass.IndicatorOrBuilder> 
           getIndicatorsOrBuilderList() {
        if (indicatorsBuilder_ != null) {
          return indicatorsBuilder_.getMessageOrBuilderList();
        } else {
          return java.util.Collections.unmodifiableList(indicators_);
        }
      }
      /**
       * <code>repeated .mail.Indicator indicators = 4;</code>
       */
      public mail.IndicatorOuterClass.Indicator.Builder addIndicatorsBuilder() {
        return getIndicatorsFieldBuilder().addBuilder(
            mail.IndicatorOuterClass.Indicator.getDefaultInstance());
      }
      /**
       * <code>repeated .mail.Indicator indicators = 4;</code>
       */
      public mail.IndicatorOuterClass.Indicator.Builder addIndicatorsBuilder(
          int index) {
        return getIndicatorsFieldBuilder().addBuilder(
            index, mail.IndicatorOuterClass.Indicator.getDefaultInstance());
      }
      /**
       * <code>repeated .mail.Indicator indicators = 4;</code>
       */
      public java.util.List<mail.IndicatorOuterClass.Indicator.Builder> 
           getIndicatorsBuilderList() {
        return getIndicatorsFieldBuilder().getBuilderList();
      }
      private com.google.protobuf.RepeatedFieldBuilderV3<
          mail.IndicatorOuterClass.Indicator, mail.IndicatorOuterClass.Indicator.Builder, mail.IndicatorOuterClass.IndicatorOrBuilder> 
          getIndicatorsFieldBuilder() {
        if (indicatorsBuilder_ == null) {
          indicatorsBuilder_ = new com.google.protobuf.RepeatedFieldBuilderV3<
              mail.IndicatorOuterClass.Indicator, mail.IndicatorOuterClass.Indicator.Builder, mail.IndicatorOuterClass.IndicatorOrBuilder>(
                  indicators_,
                  ((bitField0_ & 0x00000001) != 0),
                  getParentForChildren(),
                  isClean());
          indicators_ = null;
        }
        return indicatorsBuilder_;
      }
      @java.lang.Override
      public final Builder setUnknownFields(
          final com.google.protobuf.UnknownFieldSet unknownFields) {
        return super.setUnknownFields(unknownFields);
      }

      @java.lang.Override
      public final Builder mergeUnknownFields(
          final com.google.protobuf.UnknownFieldSet unknownFields) {
        return super.mergeUnknownFields(unknownFields);
      }


      // @@protoc_insertion_point(builder_scope:mail.MailTransaction)
    }

    // @@protoc_insertion_point(class_scope:mail.MailTransaction)
    private static final mail.MailTransactionOuterClass.MailTransaction DEFAULT_INSTANCE;
    static {
      DEFAULT_INSTANCE = new mail.MailTransactionOuterClass.MailTransaction();
    }

    public static mail.MailTransactionOuterClass.MailTransaction getDefaultInstance() {
      return DEFAULT_INSTANCE;
    }

    private static final com.google.protobuf.Parser<MailTransaction>
        PARSER = new com.google.protobuf.AbstractParser<MailTransaction>() {
      @java.lang.Override
      public MailTransaction parsePartialFrom(
          com.google.protobuf.CodedInputStream input,
          com.google.protobuf.ExtensionRegistryLite extensionRegistry)
          throws com.google.protobuf.InvalidProtocolBufferException {
        return new MailTransaction(input, extensionRegistry);
      }
    };

    public static com.google.protobuf.Parser<MailTransaction> parser() {
      return PARSER;
    }

    @java.lang.Override
    public com.google.protobuf.Parser<MailTransaction> getParserForType() {
      return PARSER;
    }

    @java.lang.Override
    public mail.MailTransactionOuterClass.MailTransaction getDefaultInstanceForType() {
      return DEFAULT_INSTANCE;
    }

  }

  private static final com.google.protobuf.Descriptors.Descriptor
    internal_static_mail_MailTransaction_descriptor;
  private static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_mail_MailTransaction_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\033mail/mail_transaction.proto\022\004mail\032\024mai" +
      "l/indicator.proto\"Y\n\017MailTransaction\022\n\n\002" +
      "id\030\001 \001(\t\022\025\n\rtx_created_at\030\003 \001(\003\022#\n\nindic" +
      "ators\030\004 \003(\0132\017.mail.Indicatorb\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          mail.IndicatorOuterClass.getDescriptor(),
        });
    internal_static_mail_MailTransaction_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_mail_MailTransaction_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mail_MailTransaction_descriptor,
        new java.lang.String[] { "Id", "TxCreatedAt", "Indicators", });
    mail.IndicatorOuterClass.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
