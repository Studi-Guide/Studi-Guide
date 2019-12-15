using System.Linq;
using NUnit.Framework;
using StudiGuideAppService.Controllers;
using Microsoft.Extensions.Logging;
using Microsoft.Extensions.Logging.Console;

namespace StudiGuideAppService.Test
{
    [TestFixture]
    public class WeatherForecastControllerTest
    {
        /// <summary>
        /// Optional
        /// </summary>
        [SetUp]
        public void Setup()
        {
        }

        [Test]
        public void CreateAndCheckWeatherController()
        {

            using (var loggerFactory = LoggerFactory.Create(builder => builder.AddConsole()))   // Need to use "using" in order to flush Console output
            {
                var logger = loggerFactory.CreateLogger<WeatherForecastController>();
                var controller = new WeatherForecastController(logger);
                var information = controller.Get();
                Assert.IsTrue(information.Any());
            }
        }
    }
}