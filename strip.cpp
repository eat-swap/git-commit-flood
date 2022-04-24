#include <cstdio>
#include <filesystem>
#include <string>
#include <cstring>

int main() {
	// In generated code, one line should not exceed 1023 chars.
	char line[1024];

	for (const auto& f : std::filesystem::recursive_directory_iterator("output")) {
		if (f.is_directory())
			continue;
		auto fileName = f.path().string();

		FILE* fIn = std::fopen(fileName.c_str(), "rb");

		if (!fIn) {
			std::printf("Cannot open to read: %s\n", fileName.c_str());
			continue;
		}

		std::string content;

		for (; std::fscanf(fIn, "%[^\n]", line) > 0; std::fscanf(fIn, "%[\n]", line)) {
			if (std::strstr(line, "this file cannot be executed directly") || std::strlen(line) == 0)
				continue;
			content += line;
			content.push_back('\n');
		}

		std::fclose(fIn);

		FILE* fOut = std::fopen(fileName.c_str(), "wb");
		auto written = std::fwrite((const void*)(content.c_str()), sizeof(char), content.length(), fOut);
		std::printf("%s expected to write %zd bytes, %zd bytes written\n", fileName.c_str(), content.length(), written);
		std::fclose(fOut);
	}
	return 0;
}
