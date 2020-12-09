defmodule Day1 do
  def main do
    File.read!("input_test")
      |> String.split("\n")
      |> Enum.map(fn x -> String.to_integer(x) end)
      |> Enum.reduce(1, fn x, acc -> x * acc end)
  end
end

IO.puts Day1.main()
