using System;
using System.IO;
using System.Text;

class MMoMemorySteam : MemoryStream
{

	public MMoMemorySteam()
	{

	}

	public MMoMemorySteam(byte[] buffer) : base(buffer)
	{

	}

	#region Short
	/// <summary>
	///  读取int
	/// </summary>
	/// <returns></returns>
	public short ReadShort()
	{
		byte[] arr = new byte[2];
		base.Read(arr, 0, 2);
		return BitConverter.ToInt16(arr, 0);
	}
	/// <summary>
	/// 写入Short
	/// </summary>
	/// <param name="value"></param>
	public void WriteShort(short value)
	{
		byte[] arr = BitConverter.GetBytes(value);
		base.Write(arr, 0, arr.Length);
	}
	#endregion

	#region UShort
	/// <summary>
	///  读取int
	/// </summary>
	/// <returns></returns>
	public ushort ReadUShort()
	{
		byte[] arr = new byte[2];
		base.Read(arr, 0, 2);
		return BitConverter.ToUInt16(arr, 0);
	}
	/// <summary>
	/// 写入UShort
	/// </summary>
	/// <param name="value"></param>
	public void WriteUShort(ushort value)
	{
		byte[] arr = BitConverter.GetBytes(value);
		base.Write(arr, 0, arr.Length);
	}
	#endregion

	#region Int
	/// <summary>
	///  读取int
	/// </summary>
	/// <returns></returns>
	public int ReadInt()
	{
		byte[] arr = new byte[4];
		base.Read(arr, 0, 4);
		return BitConverter.ToInt32(arr, 0);
	}
	/// <summary>
	/// 写入int
	/// </summary>
	/// <param name="value"></param>
	public void WriteInt(int value)
	{
		byte[] arr = BitConverter.GetBytes(value);
		base.Write(arr, 0, arr.Length);
	}
	#endregion

	#region UInt
	/// <summary>
	///  读取int
	/// </summary>
	/// <returns></returns>
	public uint ReadUInt()
	{
		byte[] arr = new byte[4];
		base.Read(arr, 0, 4);
		return BitConverter.ToUInt32(arr, 0);
	}
	/// <summary>
	/// 写入int
	/// </summary>
	/// <param name="value"></param>
	public void WriteUInt(int value)
	{
		byte[] arr = BitConverter.GetBytes(value);
		base.Write(arr, 0, arr.Length);
	}
	#endregion

	#region Long
	/// <summary>
	///  读取long
	/// </summary>
	/// <returns></returns>
	public long ReadLong()
	{
		byte[] arr = new byte[8];
		base.Read(arr, 0, 8);
		return BitConverter.ToInt64(arr, 0);
	}
	/// <summary>
	/// 写入long
	/// </summary>
	/// <param name="value"></param>
	public void WriteLong(long value)
	{
		byte[] arr = BitConverter.GetBytes(value);
		base.Write(arr, 0, arr.Length);
	}
	#endregion

	#region ULong
	/// <summary>
	///  读取ulong
	/// </summary>
	/// <returns></returns>
	public ulong ReadULong()
	{
		byte[] arr = new byte[8];
		base.Read(arr, 0, 8);
		return BitConverter.ToUInt64(arr, 0);
	}
	/// <summary>
	/// 写入ulong
	/// </summary>
	/// <param name="value"></param>
	public void WriteULong(ulong value)
	{
		byte[] arr = BitConverter.GetBytes(value);
		base.Write(arr, 0, arr.Length);
	}
	#endregion

	#region Float
	/// <summary>
	///  读取float
	/// </summary>
	/// <returns></returns>
	public float ReadFloat()
	{
		byte[] arr = new byte[4];
		base.Read(arr, 0, 4);
		return BitConverter.ToSingle(arr, 0);
	}
	/// <summary>
	/// 写入float
	/// </summary>
	/// <param name="value"></param>
	public void WriteFloat(float value)
	{
		byte[] arr = BitConverter.GetBytes(value);
		base.Write(arr, 0, arr.Length);
	}
	#endregion

	#region Double
	/// <summary>
	///  读取double
	/// </summary>
	/// <returns></returns>
	public double ReadDouble()
	{
		byte[] arr = new byte[8];
		base.Read(arr, 0, 8);
		return BitConverter.ToDouble(arr, 0);
	}
	/// <summary>
	/// 写入double
	/// </summary>
	/// <param name="value"></param>
	public void WriteDouble(double value)
	{
		byte[] arr = BitConverter.GetBytes(value);
		base.Write(arr, 0, arr.Length);
	}
	#endregion

        #region Bool
        /// <summary>
        ///  读取bool
        /// </summary>
        /// <returns></returns>
        public bool ReadBool()
        {
            byte[] arr = new byte[1];   
            return base.Read(arr, 0, 1) == 1;
        }
        /// <summary>
        /// 写入bool
        /// </summary>
        /// <param name="value"></param>
        public void WriteBool(bool value)
        {
            base.WriteByte((byte)(value == true ? 1 : 0));
        }
        #endregion

        #region Bool
        /// <summary>
        ///  读取string
        /// </summary>
        /// <returns></returns>
        public string ReadString()
        {
            uint len = ReadUInt();
            byte[] arr = new byte[len] ;
            base.Read(arr, 0, (int)len);
            return Encoding.UTF8.GetString(arr);
        }
        /// <summary>
        /// 写入string
        /// </summary>
        /// <param name="value"></param>
        public void WriteString(string value)
        {
            byte[] arr = Encoding.UTF8.GetBytes(value);
            if (arr.Length > 65535)
            {
                throw new InvalidCastException("字符串超出范围");
            }
            WriteUInt(arr.Length);
            base.Write(arr, 0, arr.Length);
        }
        #endregion

}
